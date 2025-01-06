package janusapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"

	"github.com/xprgv/janusapi-go/pkg/cmap"
	"github.com/xprgv/janusapi-go/pkg/errs"
	"github.com/xprgv/janusapi-go/pkg/logger"
	"github.com/xprgv/janusapi-go/pkg/model"
	"github.com/xprgv/janusapi-go/pkg/plugin"
	"github.com/xprgv/janusapi-go/pkg/transaction"
)

const (
	janusProtocol      = "janus-protocol"
	janusAdminProtocol = "janus-admin-protocol"
)

type Config struct {
	WebsocketUrl string
	Admin        bool
	AdminSecret  string
	AdminKey     string
}

const (
	readInnerChanSize = 10
	writeTxChanSize   = 100
)

type Janus struct {
	config  Config
	options Options

	logger logger.Logger

	conn *websocket.Conn

	ctx    context.Context
	cancel context.CancelFunc

	isConnected   atomic.Bool
	needReconnect atomic.Bool

	writeTxChan   chan *transaction.Transaction
	readInnerCh   chan []byte
	reconnectHook chan struct{}

	txNum        atomic.Uint64
	transactions cmap.ConcurrentMap[string, transaction.Transaction]

	cbFuncSessionTimeout func(sessionId uint64)
	cbFuncStreamStarted  func(sessionId uint64)
	cbFuncIceCandidate   func(sessionId uint64, iceCandidate map[string]interface{})

	cbFuncWebrtcup  func(sessionId uint64)
	cbFuncMedia     func(sessionId uint64)
	cbFuncSlowlink  func(sessionId uint64)
	cbFuncHangup    func(sessionId uint64)
	cbFuncOnCleanup func(sessionId uint64)
}

func NewJanusWebsocketApi(config Config, opts ...Option) (JanusApi, error) {
	options := getDefaultOptions()

	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, fmt.Errorf("failed to apply option: %s", err.Error())
		}
	}

	janusApi := Janus{
		config:  config,
		options: *options,

		logger: options.Logger,

		transactions: cmap.NewConcurrentMap[string, transaction.Transaction](),

		isConnected:   atomic.Bool{},
		needReconnect: atomic.Bool{},

		readInnerCh: make(chan []byte, readInnerChanSize),

		writeTxChan: make(chan *transaction.Transaction, writeTxChanSize),

		reconnectHook: make(chan struct{}, 1),
	}

	return &janusApi, nil
}

func NewJanusWebsocketApiAdmin(config Config, opts ...Option) (JanusApiAdmin, error) {
	options := getDefaultOptions()

	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, fmt.Errorf("failed to apply option: %s", err.Error())
		}
	}

	janusApi := Janus{
		config:  config,
		options: *options,

		logger: options.Logger,

		transactions: cmap.NewConcurrentMap[string, transaction.Transaction](),

		isConnected:   atomic.Bool{},
		needReconnect: atomic.Bool{},

		readInnerCh: make(chan []byte, readInnerChanSize),

		writeTxChan: make(chan *transaction.Transaction, writeTxChanSize),

		reconnectHook: make(chan struct{}, 1),
	}

	return &janusApi, nil
}

func (j *Janus) Connect(ctx context.Context) error {
	if err := j.connect(ctx); err != nil {
		return err
	}

	j.ctx, j.cancel = context.WithCancel(context.Background())
	go j.run()

	return nil
}

func (j *Janus) Close(ctx context.Context) error {
	j.isConnected.Store(false)
	j.needReconnect.Store(false)

	j.cancel() // cancel started goroutine

	if j.conn == nil {
		return nil
	}

	return j.conn.Close()
}

func (j *Janus) IsReady(ctx context.Context) bool {
	return j.isConnected.Load()
}

func (j *Janus) reconnect() {
	if !j.needReconnect.Load() {
		return
	}

	attemptCount := 0
	j.isConnected.Store(false)

	for {
		attemptCount++

		ctx, cancel := context.WithTimeout(context.Background(), j.options.JanusConnectTimeout)

		if err := j.connect(ctx); err != nil {
			cancel()
			j.logger.Error(
				fmt.Sprintf("Janus: failed to reconnect. error=[%s] attempt_count=%d", err.Error(), attemptCount),
			)
			time.Sleep(j.options.JanusReconnectInterval)
		} else {
			cancel()
			j.logger.Info(
				fmt.Sprintf("Janus: successfully reconnected. url=%s", j.config.WebsocketUrl),
			)
			if j.options.EventChannel != nil {
				select {
				case j.options.EventChannel <- model.ReconnectEvent:
				case <-time.After(500 * time.Millisecond): // skip
				}
			}
			break
		}
	}
}

func (j *Janus) run() {
	txTimeoutTicker := time.NewTicker(2 * time.Second)
	defer txTimeoutTicker.Stop()

	for {
		select {
		case <-txTimeoutTicker.C:
			j.checkOutdatedTransactions()

		case tx := <-j.writeTxChan:
			if !j.isConnected.Load() {
				tx.Cb.Err(errors.New("janus not connected"))
				continue
			}
			bin, err := json.MarshalIndent(tx.Request, "", "    ")
			if err != nil {
				tx.Cb.Err(err)
				continue
			}
			j.logger.Debug(fmt.Sprintf("to Janus >>: %s", string(bin)))
			if err := j.conn.WriteMessage(websocket.TextMessage, bin); err != nil {
				tx.Cb.Err(err)
				j.logger.Warn(fmt.Sprintf("Janus: failed to write message. error=[%s]", err.Error()))
			}

		case message := <-j.readInnerCh:
			j.processMessage(message)

		case <-j.reconnectHook:
			j.reconnect()

		case <-j.ctx.Done():
			return
		}
	}
}

func (j *Janus) connect(ctx context.Context) error {
	if j.conn != nil {
		if err := j.conn.Close(); j.isConnected.Load() {
			j.logger.Warn(fmt.Sprintf("Janus: failed to close connection. error=[%s]", err.Error()))
		}
	}
	j.isConnected.Store(false)

	protocol := janusProtocol
	if j.config.Admin {
		protocol = janusAdminProtocol
	}

	header := http.Header{"Sec-WebSocket-Protocol": {protocol}}
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, j.config.WebsocketUrl, header)
	if err != nil {
		return err
	}

	j.logger.Debug("Janus: connected")
	j.conn = conn
	j.isConnected.Store(true)
	j.needReconnect.Store(true)

	go func() {
		var CloseError *websocket.CloseError
		for {
			if _, message, err := j.conn.ReadMessage(); err == nil {
				j.readInnerCh <- message
			} else {
				if !j.isConnected.Load() && !j.needReconnect.Load() { // normal close
					return
				}
				if errors.As(err, &CloseError) || errors.Is(err, websocket.ErrCloseSent) || errors.Is(err, websocket.ErrBadHandshake) {
					j.isConnected.Store(false)
					j.reconnectHook <- struct{}{}
					return
				} else {
					j.logger.Warn(fmt.Sprintf("Janus: failed to read message. error=[%s]", err.Error()))
				}
			}
		}
	}()

	return nil
}

func (j *Janus) createTransaction(transactionType string) transaction.Transaction {
	j.txNum.Add(1)
	txId := strconv.FormatUint(j.txNum.Load(), 10)

	trans := transaction.Transaction{
		Id:   txId,
		Type: transactionType,
		Cb: &transaction.Callback{
			DoneChan: make(chan *transaction.Callback, 1),
		},
		CreatedAt: time.Now(),
	}

	j.transactions.Set(txId, trans)

	return trans
}

func (j *Janus) checkOutdatedTransactions() {
	deleteCandidates := make([]string, 0)

	j.transactions.Iter(func(id string, tx transaction.Transaction) {
		if tx.Cb.IsDone() {
			deleteCandidates = append(deleteCandidates, tx.Id)
			return
		}

		if time.Since(tx.CreatedAt) > j.options.TransactionTTL {
			tx.Cb.Err(errs.ErrTransactionTimeout)
			deleteCandidates = append(deleteCandidates, tx.Id)
		}
	})

	j.transactions.MDel(deleteCandidates)
}

func (j *Janus) processMessage(msg []byte) {
	msgString := string(msg)
	janusResponse := model.JanusResponse{}

	if err := json.Unmarshal(msg, &janusResponse); err != nil {
		j.logger.Error(fmt.Sprintf("Janus: error parse base message. error=[%s] raw_msg=[%s]", err.Error(), msgString))
		return
	}

	if janusResponse.Transaction == "" {
		j.handleEvent(janusResponse, msg)
		return
	}

	j.handleTransaction(janusResponse, msg)
}

func (j *Janus) handleEvent(janusResponse model.JanusResponse, msg []byte) {
	switch janusResponse.Janus {
	case model.JanusMessageTimeout:
		if j.cbFuncSessionTimeout != nil {
			j.cbFuncSessionTimeout(janusResponse.SessionId)
		}
	case model.JanusMessageWebrtcUp:
		if j.cbFuncWebrtcup != nil {
			j.cbFuncWebrtcup(janusResponse.SessionId)
		}
	case model.JanusMessageMedia:
		if j.cbFuncMedia != nil {
			j.cbFuncMedia(janusResponse.SessionId)
		}
	case model.JanusMessageSlowLink:
		if j.cbFuncSlowlink != nil {
			j.cbFuncSlowlink(janusResponse.SessionId)
		}
	case model.JanusMessageHangup:
		if j.cbFuncHangup != nil {
			j.cbFuncHangup(janusResponse.SessionId)
		}
	case model.JanusMessageOnCleanUp:
		if j.cbFuncOnCleanup != nil {
			j.cbFuncOnCleanup(janusResponse.SessionId)
		}
	case model.JanusMessageEvent:
		handleEventMessage(j, janusResponse, msg)
	default:
		j.logger.Warn(fmt.Sprintf("Janus: unhandled event [%s]", janusResponse.Janus))
	}
}

func handleEventMessage(j *Janus, janusResponse model.JanusResponse, msg []byte) {
	janusEvent := model.JanusResponseEvent{}
	if err := json.Unmarshal(msg, &janusEvent); err != nil {
		return
	}

	if janusEvent.PluginData.Plugin != plugin.JanusPluginStreaming {
		j.logger.Warn(fmt.Sprintf("Janus: event handlers for plugin [%s] not implemented", janusEvent.PluginData.Plugin))
		return
	}

	streamingStr, ok := janusEvent.PluginData.Data["streaming"].(string)
	if !ok {
		return
	}

	if streamingStr == model.JanusMessageEvent {
		resultMap, ok := janusEvent.PluginData.Data["result"].(map[string]interface{})
		if !ok {
			return
		}
		if resultMap["status"] == "started" {
			if j.cbFuncStreamStarted != nil {
				j.cbFuncStreamStarted(janusResponse.SessionId)
			}
		}
	}
}

func (j *Janus) handleTransaction(janusResponse model.JanusResponse, msg []byte) {
	tx, ok := j.transactions.Get(janusResponse.Transaction)
	if !ok {
		j.logger.Warn(fmt.Sprintf("Janus: transaction not found. tx_id=[%s]", janusResponse.Transaction))
		return
	}

	txHandler, ok := transaction.Handlers[tx.Type]
	if !ok {
		j.logger.Warn(fmt.Sprintf("Janus: handler for transaction [%s] not specified", tx.Type))
		return
	}

	j.logger.Debug("Janus: handle tx [%s]", tx.Type)
	txHandler(janusResponse, tx, msg)
}
