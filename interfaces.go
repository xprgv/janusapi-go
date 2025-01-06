package janusapi

import (
	"context"

	"github.com/xprgv/janusapi-go/pkg/model"
)

type JanusApi interface {
	Connect(ctx context.Context) (err error)
	Close(ctx context.Context) (err error)

	IsReady(ctx context.Context) (ok bool)

	ServerInfo(ctx context.Context) (serverInfo model.ServerInfo, err error)

	MountpointList(ctx context.Context, sessionId uint64, handleId uint64) (list []model.MountpointShort, err error)
	MountpointInfo(
		ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64, mountpointSecret string,
	) (mountpoint model.Mountpoint, err error)
	MountpointCreate(ctx context.Context, sessionId uint64, handleId uint64, mountpoint model.MountpointLegacy) (err error)
	MountpointDestroy(
		ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64, mountpointSecret string,
	) (err error)

	SessionCreate(ctx context.Context, opaqueId string) (sessionId uint64, err error)
	SessionAttach(ctx context.Context, sessionId uint64, pluginName string) (handleId uint64, err error)
	SessionKeepalive(ctx context.Context, sessionId uint64) (err error)
	SessionDestroy(ctx context.Context, sessionId uint64) (err error)

	StreamWatch(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64) (sdpOffer string, err error)
	StreamStart(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64, sdpAnswer string) (err error)
	StreamPause(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64) (err error)
	StreamResume(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64) (err error)
	StreamSwitch(ctx context.Context, sessionId uint64, handleId uint64, mountpointId uint64) (err error)
	IceCandidate(ctx context.Context, sessionId uint64, handleId uint64, iceCandidate map[string]interface{}) (err error)

	OnSessionTimeout(cb func(sessionId uint64))
	OnStreamStarted(cb func(sessionId uint64))
	OnIceCandidate(cb func(sessionId uint64, iceCandidate map[string]interface{}))
	OnWebrtcUp(cb func(sessionId uint64))
	OnMedia(cb func(sessionId uint64))
	OnSlowLink(cb func(sessionId uint64))
	OnHangup(cb func(sessionId uint64))
	OnCleanup(cb func(sessionId uint64))
}

type JanusApiAdmin interface {
	Connect(ctx context.Context) (err error)
	Close(ctx context.Context) (err error)
	IsReady(ctx context.Context) (ok bool)
	HandleInfo(ctx context.Context, sessionId uint64, handleId uint64) (handleInfo model.HandleInfo, err error)
}
