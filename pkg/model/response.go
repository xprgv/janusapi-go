package model

type JanusResponse struct {
	Janus       string `json:"janus"`
	SessionId   uint64 `json:"session_id"`
	Sender      uint64 `json:"sender"`
	Transaction string `json:"Transaction"`
}

type JanusResponseMessage struct {
	Janus       string                   `json:"janus"`
	SessionId   uint64                   `json:"session_id"`
	Sender      uint64                   `json:"sender"`
	Transaction string                   `json:"Transaction"`
	Jsep        Jsep                     `json:"jsep"`
	Body        JanusResponseMessageBody `json:"body"`
}

type JanusResponseMessageBody struct {
	Request string `json:"request"`
}

type JanusResponseEvent struct {
	Janus       string     `json:"janus"`
	SessionId   uint64     `json:"session_id"`
	Sender      uint64     `json:"sender"`
	Transaction string     `json:"Transaction"`
	PluginData  PluginData `json:"plugindata"`
}

type JanusResponseSuccess struct {
	Janus       string      `json:"janus"`
	Transaction string      `json:"Transaction"`
	Data        SuccessData `json:"data"`
	PluginData  PluginData  `json:"plugindata"`
}

type SuccessData struct {
	Id uint64 `json:"id"`
}

type JanusResponseError struct {
	Janus       string    `json:"janus"`
	Transaction string    `json:"Transaction"`
	SessionId   uint64    `json:"session_id"`
	ErrorData   ErrorData `json:"error"`
}

type ErrorData struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

type JanusResponseHandleInfo struct {
	Janus       string     `json:"janus"`
	SessionId   uint64     `json:"session_id"`
	HandleId    uint64     `json:"handle_id"`
	Transaction string     `json:"Transaction"`
	Info        HandleInfo `json:"info"`
}
