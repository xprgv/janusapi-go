package janusapi

import (
	"context"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/model"
	"github.com/xprgv/janusapi-go/pkg/transaction"
)

func (j *Janus) ServerInfo(ctx context.Context) (model.ServerInfo, error) {
	tx := j.createTransaction(transaction.TxTypeInfo)
	tx.Request = model.JanusRequest{
		Janus:       tx.Type,
		Transaction: tx.Id,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		if result.Error != nil {
			return model.ServerInfo{}, result.Error
		}
		serverInfo, ok := tx.Cb.Reply.(model.ServerInfo)
		if !ok {
			return model.ServerInfo{}, fmt.Errorf("can't convert reply to server info")
		}
		return serverInfo, nil
	case <-ctx.Done():
		return model.ServerInfo{}, ctx.Err()
	}
}
