package model

type JanusRequest struct {
	Janus       string `json:"janus"`
	Transaction string `json:"transaction"`
}

type JanusRequestCreate struct {
	Janus       string `json:"janus"`
	Transaction string `json:"transaction"`
}

type JanusRequestAttach struct {
	Janus       string `json:"janus"`
	Plugin      string `json:"plugin"`
	SessionId   uint64 `json:"session_id"`
	Transaction string `json:"transaction"`
}

type JanusRequestKeepalive struct {
	Janus       string `json:"janus"`
	SessionId   uint64 `json:"session_id"`
	Transaction string `json:"transaction"`
}

type JanusRequestDestroy struct {
	Janus       string `json:"janus"`
	SessionId   uint64 `json:"session_id"`
	Transaction string `json:"transaction"`
}

type JanusRequestListMountpoint struct {
	Janus       string                         `json:"janus"`
	Body        JanusRequestListMountpointBody `json:"body"`
	SessionId   uint64                         `json:"session_id"`
	Transaction string                         `json:"transaction"`
	HandleId    uint64                         `json:"handle_id"`
}

type JanusRequestListMountpointBody struct {
	Request string `json:"request"`
}

type JanusRequestInfoMountpoint struct {
	Janus       string                         `json:"janus"`
	Body        JanusRequestInfoMountpointBody `json:"body"`
	Transaction string                         `json:"transaction"`
	SessionId   uint64                         `json:"session_id"`
	HandleId    uint64                         `json:"handle_id"`
}

type JanusRequestInfoMountpointBody struct {
	Request string `json:"request"`
	Id      uint64 `json:"id"`
	Secret  string `json:"secret"`
}

type JanusRequestCreateMountpoint struct {
	Janus       string                           `json:"janus"`
	Body        JanusRequestCreateMountpointBody `json:"body"`
	SessionId   uint64                           `json:"session_id"`
	Transaction string                           `json:"transaction"`
	HandleId    uint64                           `json:"handle_id"`
}

// TODO: fix legacy format
type JanusRequestCreateMountpointBody struct {
	Request        string `json:"request"`
	Id             uint64 `json:"id"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Metadata       string `json:"metadata"`
	Audio          bool   `json:"audio"`
	Video          bool   `json:"video"`
	Data           bool   `json:"data"`
	Videoport      int64  `json:"videoport"`
	VideoRtcpPort  int64  `json:"videortcpport,omitempty"`
	Videopt        int64  `json:"videopt"`
	Videortpmap    string `json:"videortpmap"`
	Videofmtp      string `json:"videofmtp"`
	Audioport      int64  `json:"audioport"`
	Audiortcpport  int64  `json:"audiortcpport,omitempty"`
	Dataport       int64  `json:"dataport"`
	Audiopt        int64  `json:"audiopt"`
	Audiortpmap    string `json:"audiortpmap"`
	Audiofmtp      string `json:"audiofmtp"`
	Secret         string `json:"secret"`
	AdminKey       string `json:"admin_key"`
	VideoSimulcast bool   `json:"videosimulcast"`
	Videoport2     int64  `json:"videoport2"`
	Videoport3     int64  `json:"videoport3"`
	Datatype       string `json:"datatype"`
	DataBufferMsg  bool   `json:"databuffermsg"`
	VideoBufferKf  bool   `json:"videobufferkf"`
	VideoSkew      bool   `json:"videoskew"`
	Videosvc       bool   `json:"videosvc"`
	Collision      int64  `json:"collision"`
	Threads        int64  `json:"threads"`
}

type JanusRequestDestroyMountpoint struct {
	Janus       string                            `json:"janus"`
	Body        JanusRequestDestroyMountpointBody `json:"body"`
	SessionId   uint64                            `json:"session_id"`
	Transaction string                            `json:"transaction"`
	HandleId    uint64                            `json:"handle_id"`
}

type JanusRequestDestroyMountpointBody struct {
	Id      uint64 `json:"id"`
	Secret  string `json:"secret"`
	Request string `json:"request"`
}

type JanusRequestWatch struct {
	Janus       string                `json:"janus"`
	Body        JanusRequestWatchBody `json:"body"`
	SessionId   uint64                `json:"session_id"`
	Transaction string                `json:"transaction"`
	HandleId    uint64                `json:"handle_id"`
}

type JanusRequestWatchBody struct {
	Id      uint64 `json:"id"`
	Request string `json:"request"`
}

type JanusRequestStart struct {
	Janus       string                `json:"janus"`
	Body        JanusRequestStartBody `json:"body"`
	SessionId   uint64                `json:"session_id"`
	Transaction string                `json:"transaction"`
	HandleId    uint64                `json:"handle_id"`
	Jsep        Jsep                  `json:"jsep"`
}

type JanusRequestStartBody struct {
	Id      uint64 `json:"id"`
	Request string `json:"request"`
}

type JanusRequestPause struct {
	Janus       string                `json:"janus"`
	Body        JanusRequestPauseBody `json:"body"`
	SessionId   uint64                `json:"session_id"`
	Transaction string                `json:"transaction"`
	HandleId    uint64                `json:"handle_id"`
}

type JanusRequestPauseBody struct {
	Id      uint64 `json:"id"`
	Request string `json:"request"`
}

type JanusRequestResume struct {
	Janus       string                 `json:"janus"`
	Body        JanusRequestResumeBody `json:"body"`
	SessionId   uint64                 `json:"session_id"`
	Transaction string                 `json:"transaction"`
	HandleId    uint64                 `json:"handle_id"`
}

type JanusRequestResumeBody struct {
	Id      uint64 `json:"id"`
	Request string `json:"request"`
}

type JanusRequestSwitch struct {
	Janus       string                 `json:"janus"`
	Body        JanusRequestSwitchBody `json:"body"`
	SessionId   uint64                 `json:"session_id"`
	Transaction string                 `json:"transaction"`
	HandleId    uint64                 `json:"handle_id"`
}

type JanusRequestSwitchBody struct {
	Id      uint64 `json:"id"`
	Request string `json:"request"`
}

type JanusRequestTrickle struct {
	Janus       string                 `json:"janus"`
	SessionId   uint64                 `json:"session_id"`
	Transaction string                 `json:"transaction"`
	HandleId    uint64                 `json:"handle_id"`
	Candidate   map[string]interface{} `json:"candidate"`
}

type JanusRequestHandleInfo struct {
	Janus       string `json:"janus"`
	SessionId   uint64 `json:"session_id"`
	HandleId    uint64 `json:"handle_id"`
	Transaction string `json:"transaction"`
	AdminSecret string `json:"admin_secret"`
}
