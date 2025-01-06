package model

type Jsep struct {
	Type string `json:"type"`
	Sdp  string `json:"sdp"`
}

type PluginData struct {
	Plugin string                 `json:"plugin"`
	Data   map[string]interface{} `json:"data"`
}
