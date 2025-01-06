package janusapi

import (
	"context"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/model"
	"github.com/xprgv/janusapi-go/pkg/transaction"
)

func (j *Janus) StreamWatch(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64) (string, error) {
	tx := j.createTransaction(transaction.TxTypeWatch)
	tx.Request = model.JanusRequestWatch{
		Janus: "message",
		Body: model.JanusRequestWatchBody{
			Request: tx.Type,
			Id:      mountpointId,
		},
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		if result.Error != nil {
			return "", result.Error
		}
		jsep, ok := result.Reply.(model.Jsep)
		if !ok {
			return "", fmt.Errorf("can't convert watch.reply to Jsep")
		}
		return jsep.Sdp, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (j *Janus) StreamStart(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64, sdpAnswer string) error {
	tx := j.createTransaction(transaction.TxTypeStart)
	tx.Request = model.JanusRequestStart{
		Janus: "message",
		Body: model.JanusRequestStartBody{
			Request: tx.Type,
			Id:      mountpointId,
		},
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
		Jsep: model.Jsep{
			Type: "answer",
			Sdp:  sdpAnswer,
		},
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		return result.Error
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (j *Janus) StreamPause(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64) error {
	tx := j.createTransaction(transaction.TxTypePause)
	tx.Request = model.JanusRequestPause{
		Janus: "message",
		Body: model.JanusRequestPauseBody{
			Request: tx.Type,
			Id:      mountpointId,
		},
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		return result.Error
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (j *Janus) StreamResume(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64) error {
	tx := j.createTransaction(transaction.TxTypeResume)

	tx.Request = model.JanusRequestResume{
		Janus: "message",
		Body: model.JanusRequestResumeBody{
			Id:      mountpointId,
			Request: "start",
		},
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		return result.Error
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (j *Janus) StreamSwitch(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64) error {
	tx := j.createTransaction(transaction.TxTypeSwitch)

	tx.Request = model.JanusRequestSwitch{
		Janus: "message",
		Body: model.JanusRequestSwitchBody{
			Request: tx.Type,
			Id:      mountpointId,
		},
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		return result.Error
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (j *Janus) IceCandidate(ctx context.Context, sessionId uint64, handleId uint64, iceCandidate map[string]interface{}) error {
	tx := j.createTransaction(transaction.TxTypeTrickle)

	tx.Request = model.JanusRequestTrickle{
		Janus:       tx.Type,
		Candidate:   iceCandidate,
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		return result.Error
	case <-ctx.Done():
		return ctx.Err()
	}
}
