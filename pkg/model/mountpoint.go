package model

type Mountpoint struct {
	Id          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Metadata    string  `json:"metadata"`
	Secret      string  `json:"secret"`
	Enabled     bool    `json:"enabled"`
	Viewers     int64   `json:"viewers"`
	Type        string  `json:"type"`
	Media       []Media `json:"media"`
	Collision   int64   `json:"collision"`
}

type MediaType string

const (
	DataType  MediaType = "data"
	VideoType MediaType = "video"
	AudioType MediaType = "audio"
)

type Media struct {
	Type           MediaType `json:"type,omitempty"`
	Mid            string    `json:"mid,omitempty"`
	Label          string    `json:"label,omitempty"`
	Pt             uint64    `json:"pt,omitempty"`
	Codec          string    `json:"codec,omitempty"`
	DataType       string    `json:"datatype,omitempty"`
	RtpMap         string    `json:"rtpmap,omitempty"`
	Fmtp           string    `json:"fmtp,omitempty"`
	Port           uint64    `json:"port,omitempty"`
	AgeMs          int64     `json:"age_ms,omitempty"`
	VideoSimulcast bool      `json:"videosimulcast,omitempty"`
}

type MountpointShort struct {
	Id          uint64       `json:"id"`
	Type        string       `json:"type"`
	Description string       `json:"description"`
	Enabled     bool         `json:"enabled"`
	Media       []ShortMedia `json:"media"`
}

type ShortMedia struct {
	Mid   string    `json:"mid,omitempty"`
	Type  MediaType `json:"type,omitempty"`
	Label string    `json:"label,omitempty"`
	AgeMs int64     `json:"age_ms,omitempty"`
}

type MountpointLegacy struct {
	Id             uint64 `json:"id"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Metadata       string `json:"metadata"`
	Viewers        int64  `json:"viewers"`
	Enabled        bool   `json:"enabled"`
	Audio          bool   `json:"audio"`
	Video          bool   `json:"video"`
	Data           bool   `json:"data"`
	Dataport       int64  `json:"dataport"`
	Videoport      int64  `json:"videoport"`
	VideoRtcpPort  int64  `json:"videortcpport"`
	VideoSimulcast bool   `json:"videosimulcast"`
	Videoport2     int64  `json:"videoport2"`
	Videoport3     int64  `json:"videoport3"`
	// VideoRtcpPort  int    `json:"videortcpport,omitempty"`
	Videopt     int64  `json:"videopt"`
	Videortpmap string `json:"videortpmap"`
	Videofmtp   string `json:"videofmtp"`
	VideoAge    int64  `json:"video_age_ms"`
	Audioport   int64  `json:"audioport"`
	// AudioRtcpPort int    `json:"audiortcpport,omitempty"`
	AudioRtcpPort int64  `json:"audiortcpport"`
	Audiopt       int64  `json:"audiopt"`
	Audiortpmap   string `json:"audiortpmap"`
	Audiofmtp     string `json:"audiofmtp"`
	AudioAge      int64  `json:"audio_age_ms"`
	Secret        string `json:"secret"`
	Datatype      string `json:"datatype"`
	DataBufferMsg bool   `json:"databuffermsg"`
	VideoBufferKf bool   `json:"videobufferkf"`
	VideoSkew     bool   `json:"videoskew"`
	Videosvc      bool   `json:"videosvc"`
	Collision     int64  `json:"collision"`
	Threads       int64  `json:"threads"`
}

// type MountpointShort struct {
// 	Id          int64  `json:"id"`
// 	Type        string `json:"type"`
// 	Description string `json:"description"`
// 	Enabled     bool   `json:"enabled"`
// 	AudioAgeMs  int64  `json:"audio_age_ms"`
// 	VideoAgeMs  int64  `json:"video_age_ms"`
// }
