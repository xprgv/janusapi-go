package transaction

import (
	"sync/atomic"
	"time"

	"github.com/xprgv/janusapi-go/pkg/model"
)

const (
	TxTypeInfo              string = "info"
	TxTypeCreate            string = "create"
	TxTypeAttach            string = "attach"
	TxTypeWatch             string = "watch"
	TxTypePrepare           string = "prepare"
	TxTypeStart             string = "start"
	TxTypeSwitch            string = "switch"
	TxTypeTrickle           string = "trickle"
	TxTypePause             string = "pause"
	TxTypeResume            string = "resume"
	TxTypeDestroy           string = "destroy"
	TxTypeKeepalive         string = "keepalive"
	TxTypeCreateMountpoint  string = "create_mountpoint"
	TxTypeDestroyMountpoint string = "destroy_mountpoint"
	TxTypeListMountpoint    string = "list_mountpoint"
	TxTypeInfoMountpoint    string = "info_mountpoint"
	TxTypeJoin              string = "join"
	TxTypeConfigure         string = "configure"
	TxTypeRtpForward        string = "rtp_forward"
	TxTypeHandleInfo        string = "handle_info"
	TxTypeError             string = "error"
)

type Transaction struct {
	Id        string
	Type      string
	CreatedAt time.Time
	Cb        *Callback
	Request   interface{}
}

type Callback struct {
	Reply    interface{}
	Error    error
	DoneChan chan *Callback
	doneFlag atomic.Bool
}

func (txCb *Callback) IsDone() bool {
	return txCb.doneFlag.Load()
}

func (txCb *Callback) Ok(reply interface{}) {
	txCb.Reply = reply
	txCb.doneFlag.Store(true)
	txCb.DoneChan <- txCb
}

func (txCb *Callback) Err(err error) {
	txCb.Error = err
	txCb.doneFlag.Store(true)
	txCb.DoneChan <- txCb
}

type Handler func(janusResponse model.JanusResponse, tx Transaction, msg []byte)

var Handlers = map[string]Handler{
	// server related
	TxTypeInfo: handleTxInfo,

	// session related
	TxTypeCreate:    handleTxCreate,
	TxTypeAttach:    handleTxAttach,
	TxTypeKeepalive: handleTxKeepalive,
	TxTypeDestroy:   handleTxDestroy,

	// mountpoint related
	TxTypeListMountpoint:    handleTxListMountpoint,
	TxTypeInfoMountpoint:    handleTxInfoMountpoint,
	TxTypeCreateMountpoint:  handleTxCreateMountpoint,
	TxTypeDestroyMountpoint: handleTxDestroyMountpoint,

	// stream related
	TxTypeWatch:   handleTxWatch,
	TxTypeStart:   handleTxStart,
	TxTypePause:   handleTxPause,
	TxTypeResume:  handleTxResume,
	TxTypeSwitch:  handleTxSwitch,
	TxTypeTrickle: handleTxIceCandidate,

	// available only in Janus-admin-protocol
	TxTypeHandleInfo: handleTxHandleInfo,
}
