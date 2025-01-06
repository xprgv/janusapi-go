package janusapi

import (
	"context"
	"fmt"

	"github.com/xprgv/janusapi-go/pkg/model"
	"github.com/xprgv/janusapi-go/pkg/transaction"
)

func (j *Janus) MountpointList(ctx context.Context, sessionId uint64, handleId uint64) ([]model.MountpointShort, error) {
	tx := j.createTransaction(transaction.TxTypeListMountpoint)
	tx.Request = model.JanusRequestListMountpoint{
		Janus: "message",
		Body: model.JanusRequestListMountpointBody{
			Request: "list",
		},
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		if result.Error != nil {
			return []model.MountpointShort{}, result.Error
		}
		mountpoints, ok := tx.Cb.Reply.(*[]model.MountpointShort)
		if !ok {
			return []model.MountpointShort{}, fmt.Errorf("can't convert reply to mountpoint list")
		}
		return *mountpoints, nil
	case <-ctx.Done():
		return []model.MountpointShort{}, ctx.Err()
	}
}

func (j *Janus) MountpointInfo(
	ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64, mountpointSecret string,
) (model.Mountpoint, error) {
	tx := j.createTransaction(transaction.TxTypeInfoMountpoint)
	tx.Request = model.JanusRequestInfoMountpoint{
		Janus: "message",
		Body: model.JanusRequestInfoMountpointBody{
			Request: "info",
			Id:      mountpointId,
			Secret:  mountpointSecret,
		},
		Transaction: tx.Id,
		SessionId:   sessionId,
		HandleId:    handleId,
	}

	j.writeTxChan <- &tx

	select {
	case result := <-tx.Cb.DoneChan:
		if result.Error != nil {
			return model.Mountpoint{}, result.Error
		}
		mountpoint, ok := result.Reply.(*model.Mountpoint)
		if !ok {
			return model.Mountpoint{}, fmt.Errorf("can't convert reply to mountpoint")
		}

		return *mountpoint, nil
	case <-ctx.Done():
		return model.Mountpoint{}, ctx.Err()
	}
}

// TODO: change legacy mountpoint format to the new
func (j *Janus) MountpointCreate(ctx context.Context, sessionId uint64, handleId uint64, mountpoint model.MountpointLegacy) error {
	tx := j.createTransaction(transaction.TxTypeCreateMountpoint)

	tx.Request = model.JanusRequestCreateMountpoint{
		Janus: "message",
		Body: model.JanusRequestCreateMountpointBody{
			Request:        "create",
			Id:             mountpoint.Id,
			Type:           mountpoint.Type,
			Name:           mountpoint.Name,
			Description:    mountpoint.Description,
			Metadata:       mountpoint.Metadata,
			Audio:          mountpoint.Audio,
			Video:          mountpoint.Video,
			Data:           mountpoint.Data,
			Videoport:      mountpoint.Videoport,
			VideoRtcpPort:  mountpoint.VideoRtcpPort,
			Videopt:        mountpoint.Videopt,
			Videortpmap:    mountpoint.Videortpmap,
			Videofmtp:      mountpoint.Videofmtp,
			Audioport:      mountpoint.Audioport,
			Audiortcpport:  mountpoint.AudioRtcpPort,
			Audiopt:        mountpoint.Audiopt,
			Audiortpmap:    mountpoint.Audiortpmap,
			Audiofmtp:      mountpoint.Audiofmtp,
			Dataport:       mountpoint.Dataport,
			Secret:         mountpoint.Secret,
			AdminKey:       j.config.AdminKey,
			VideoSimulcast: mountpoint.VideoSimulcast,
			Videoport2:     mountpoint.Videoport2,
			Videoport3:     mountpoint.Videoport3,
			Datatype:       mountpoint.Datatype,
			DataBufferMsg:  mountpoint.DataBufferMsg,
			VideoBufferKf:  mountpoint.VideoBufferKf,
			VideoSkew:      mountpoint.VideoSkew,
			Videosvc:       mountpoint.Videosvc,
			Collision:      mountpoint.Collision,
			Threads:        mountpoint.Threads,
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

func (j *Janus) MountpointDestroy(
	ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64, mountpointSecret string,
) error {
	tx := j.createTransaction(transaction.TxTypeDestroyMountpoint)
	tx.Request = model.JanusRequestDestroyMountpoint{
		Janus: "message",
		Body: model.JanusRequestDestroyMountpointBody{
			Request: "destroy",
			Id:      mountpointId,
			Secret:  mountpointSecret,
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
