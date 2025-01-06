package janusapi

import (
	"context"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/model"
	"github.com/xprgv/janusapi-go/pkg/transaction"
)

func (j *Janus) SessionCreate(ctx context.Context, opaqueId string) (uint64, error) {
	tx := j.createTransaction(transaction.TxTypeCreate)
	tx.Request = model.JanusRequestCreate{
		Janus:       tx.Type,
		Transaction: tx.Id,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		if result.Error != nil {
			return 0, result.Error
		}
		sessionId, ok := result.Reply.(uint64)
		if !ok {
			return 0, fmt.Errorf("can't convert create.reply to session_id")
		}
		return sessionId, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func (j *Janus) SessionAttach(ctx context.Context, sessionId uint64, pluginName string) (uint64, error) {
	tx := j.createTransaction(transaction.TxTypeAttach)
	tx.Request = model.JanusRequestAttach{
		Janus:       tx.Type,
		Plugin:      pluginName,
		Transaction: tx.Id,
		SessionId:   sessionId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		if result.Error != nil {
			return 0, result.Error
		}
		handleId, ok := result.Reply.(uint64)
		if !ok {
			return 0, fmt.Errorf("can't convert attach.reply to handle_id")
		}
		return handleId, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func (j *Janus) SessionKeepalive(ctx context.Context, sessionId uint64) error {
	tx := j.createTransaction(transaction.TxTypeKeepalive)
	tx.Request = model.JanusRequestKeepalive{
		Janus:       tx.Type,
		Transaction: tx.Id,
		SessionId:   sessionId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		return result.Error
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (j *Janus) SessionDestroy(ctx context.Context, sessionId uint64) error {
	tx := j.createTransaction(transaction.TxTypeDestroy)
	tx.Request = model.JanusRequestDestroy{
		Janus:       tx.Type,
		Transaction: tx.Id,
		SessionId:   sessionId,
	}

	j.writeTxChan <- &tx

	select {
	case <-tx.Cb.DoneChan:
		return tx.Cb.Error
	case <-ctx.Done():
		return ctx.Err()
	}
}
