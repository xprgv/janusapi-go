package janusapi

import (
	"context"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/model"
	"github.com/xprgv/janusapi-go/pkg/transaction"
)

func (j *Janus) HandleInfo(ctx context.Context, sessionId uint64, handleId uint64) (handleInfo model.HandleInfo, err error) {
	tx := j.createTransaction(transaction.TxTypeHandleInfo)

	tx.Request = model.JanusRequestHandleInfo{
		Janus:       tx.Type,
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
		AdminSecret: j.config.AdminSecret,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		if res, ok := result.Reply.(model.HandleInfo); ok {
			return res, nil
		}
		return model.HandleInfo{}, fmt.Errorf("can't convert reply to handle info")
	case <-ctx.Done():
		return model.HandleInfo{}, ctx.Err()
	}
}
