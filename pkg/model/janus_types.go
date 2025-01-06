package model

const (
	JanusMessageSuccess   = "success"
	JanusMessageError     = "error"
	JanusMessageAck       = "ack"
	JanusMessageEvent     = "event"
	JanusMessageWebrtcUp  = "webrtcup"
	JanusMessageTimeout   = "timeout"
	JanusMessageMedia     = "media"
	JanusMessageSlowLink  = "slowlink"
	JanusMessageHangup    = "hangup"
	JanusMessageOnCleanUp = "oncleanup"
)

type HandleInfo struct {
	Janus       string            `json:"janus"`
	Transaction string            `json:"transaction"`
	SessionID   int64             `json:"session_id"`
	HandleID    int64             `json:"handle_id"`
	Info        HandleInfoDetails `json:"info"`
}

type HandleInfoDetails struct {
	SessionID           int64              `json:"session_id"`
	SessionLastActivity int64              `json:"session_last_activity"`
	SessionTransport    string             `json:"session_transport"`
	HandleID            int64              `json:"handle_id"`
	OpaqueID            string             `json:"opaque_id"`
	LoopRunning         bool               `json:"loop-running"`
	Created             int64              `json:"created"`
	CurrentTime         int64              `json:"current_time"`
	Plugin              string             `json:"plugin"`
	PluginSpecific      PluginSpecificInfo `json:"plugin_specific"`
	Flags               map[string]bool    `json:"flags"`
	AgentCreated        int64              `json:"agent_created"`
	IceMode             string             `json:"ice-mode"`
	IceRole             string             `json:"ice-role"`
	SDPs                SDPInfo            `json:"sdps"`
	QueuedPackets       int                `json:"queued-packets"`
	Streams             []Stream           `json:"streams"`
}

type PluginSpecificInfo struct {
	State             string `json:"state"`
	MountpointId      int    `json:"mountpoint_id"`
	MountpointName    string `json:"mountpoint_name"`
	MountpointViewers int    `json:"mountpoint_viewers"`
}

type SDPInfo struct {
	Profile string `json:"profile"`
	Local   string `json:"local"`
	Remote  string `json:"remote"`
}

type Stream struct {
	RtcpStats  RtcpStats   `json:"rtcp_stats"`
	Components []Component `json:"components"`
}

type RtcpStats struct {
	Audio Audio `json:"audio"`
	Video Video `json:"video"`
}

type Audio struct{}
type Video struct{}

type Component struct {
	State    string   `json:"state"`
	InStats  InStats  `json:"in_stats"`
	OutStats OutStats `json:"out_stats"`
}

type InStats struct {
	AudioPackets      int `json:"audio_packets"`
	AudioBytes        int `json:"audio_bytes"`
	AudioBytesLastSec int `json:"audio_bytes_lastsec"`
	VideoPackets      int `json:"video_packets"`
	VideoBytes        int `json:"video_bytes"`
	VideoBytesLastSec int `json:"video_bytes_lastsec"`
}

type OutStats struct {
	AudioPackets      int `json:"audio_packets"`
	AudioBytes        int `json:"audio_bytes"`
	AudioBytesLastSec int `json:"audio_bytes_lastsec"`
	VideoPackets      int `json:"video_packets"`
	VideoBytes        int `json:"video_bytes"`
	VideoBytesLastSec int `json:"video_bytes_lastsec"`
}

type ServerInfo struct {
	Janus                 string                   `json:"server_info"`
	Transaction           string                   `json:"transaction"`
	Name                  string                   `json:"name"`
	Version               int                      `json:"version"`
	VersionString         string                   `json:"version_string"`
	Author                string                   `json:"author"`
	CommitHash            string                   `json:"commit-hash"`
	CompileTime           string                   `json:"compile-time"`
	LogToStdout           bool                     `json:"log-to-stdout"`
	LogToFile             bool                     `json:"log-to-file"`
	DataChannels          bool                     `json:"data_channels"`
	AcceptingNewSessions  bool                     `json:"accepting-new-sessions"`
	SessionTimeout        int                      `json:"session-timeout"`
	ReclaimSessionTimeout int                      `json:"reclaim-session-timeout"`
	CandidatesTimeout     int                      `json:"candidates-timeout"`
	ServerName            string                   `json:"server-name"`
	LocalIp               string                   `json:"local-ip"`
	IpV6                  bool                     `json:"ipv6"`
	IceLite               bool                     `json:"ice-lite"`
	IceTcp                bool                     `json:"ice-tcp"`
	IceNomination         string                   `json:"ice-nomination"`
	IceKeepaliveConncheck bool                     `json:"ice-keepalive-conncheck"`
	FullTrickle           bool                     `json:"full-trickle"`
	MdnsEnabled           bool                     `json:"mdns-enabled"`
	MinNackQueue          int                      `json:"min-nack-queue"`
	NackOptimizations     bool                     `json:"nack-optimizations"`
	TwccPeriod            int                      `json:"twcc-period"`
	DtlsMtu               int                      `json:"dtls-mtu"`
	StaticEventLoops      int                      `json:"static-event-loops"`
	ApiSecret             bool                     `json:"api_secret"`
	AuthToken             bool                     `json:"auth_token"`
	EventHandlers         bool                     `json:"event_handlers"`
	OpaqueidInApi         bool                     `json:"opaqueid_in_api"`
	Dependencies          map[string]string        `json:"dependencies"`
	Transports            map[string]TransportInfo `json:"transports"`
	Events                map[string]interface{}   `json:"events"`
	Loggers               map[string]interface{}   `json:"loggers"`
	Plugins               map[string]PluginInfo    `json:"plugins"`
}

type TransportInfo struct {
	Name          string `json:"name"`
	Author        string `json:"author"`
	Description   string `json:"description"`
	VersionString string `json:"version_string"`
	Version       int    `json:"version"`
}

type PluginInfo struct {
	Name          string `json:"name"`
	Author        string `json:"author"`
	Description   string `json:"description"`
	VersionString string `json:"version_string"`
	Version       int    `json:"version"`
}

type PluginDataStreamingStream struct {
	Id          uint64 `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	IsPrivate   bool   `json:"is_private"`
	AudioPort   uint64 `json:"audio_port"`
	VideoPort   uint64 `json:"video_port"`
}
