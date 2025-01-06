package plugin

const (
	JanusPluginStreaming  string = "janus.plugin.streaming"
	JanusPluginVideoCall  string = "janus.plugin.videocall"
	JanusPluginRecordPlay string = "janus.plugin.recordplay"
	JanusPluginNoSip      string = "janus.plugin.nosip"
	JanusPluginTextRoom   string = "janus.plugin.textroom"
	JanusPluginVideoRoom  string = "janus.plugin.videoroom"
	JanusPluginEchoTest   string = "janus.plugin.echotest"
)

/* Error codes */
const (
	JanusStreamingErrorNoMessage        = 450
	JanusStreamingErrorInvalidJson      = 451
	JanusStreamingErrorInvalidRequest   = 452
	JanusStreamingErrorMissingElement   = 453
	JanusStreamingErrorInvalidElement   = 454
	JanusStreamingErrorNoSuchMountpoint = 455
	JanusStreamingErrorCantCreate       = 456
	JanusStreamingErrorUnauthorized     = 457
	JanusStreamingErrorCantSwitch       = 458
	JanusStreamingErrorCantRecord       = 459
	JanusStreamingErrorInvalidState     = 460
	JanusStreamingErrorUnknownError     = 470
)
