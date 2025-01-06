package janusapi

func (j *Janus) OnSessionTimeout(cb func(sessionId uint64)) {
	j.cbFuncSessionTimeout = cb
}

func (j *Janus) OnStreamStarted(cb func(sessionId uint64)) {
	j.cbFuncStreamStarted = cb
}

func (j *Janus) OnIceCandidate(cb func(sessionId uint64, iceCandidate map[string]interface{})) {
	j.cbFuncIceCandidate = cb
}

func (j *Janus) OnWebrtcUp(cb func(sessionId uint64)) {
	j.cbFuncWebrtcup = cb
}

func (j *Janus) OnMedia(cb func(sessionId uint64)) {
	j.cbFuncMedia = cb
}

func (j *Janus) OnSlowLink(cb func(sessionId uint64)) {
	j.cbFuncSlowlink = cb
}

func (j *Janus) OnHangup(cb func(sessionId uint64)) {
	j.cbFuncHangup = cb
}

func (j *Janus) OnCleanup(cb func(sessionId uint64)) {
	j.cbFuncOnCleanup = cb
}
