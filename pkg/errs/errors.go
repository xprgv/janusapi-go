package errs

import "errors"

var (
	ErrTransactionTimeout = errors.New("transaction timeout")
	ErrJanusNotConnected  = errors.New("janus not connected")
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

const (
	// JanusErrorUnauthorized /*! \brief Unauthorized (can only happen when using apisecret/auth token) */
	JanusErrorUnauthorized = 403

	// JanusErrorUnauthorizedPlugin /*! \brief Unauthorized access to a plugin (can only happen when using auth token) */
	JanusErrorUnauthorizedPlugin = 405

	// JanusErrorUnknown /*! \brief Unknown/undocumented error */
	JanusErrorUnknown = 490

	// JanusErrorTransportSpecific /*! \brief Transport related error */
	JanusErrorTransportSpecific = 450

	// JanusErrorMissingRequest /*! \brief The request is missing in the message */
	JanusErrorMissingRequest = 452

	// JanusErrorUnknownRequest /*! \brief The Janus core does not suppurt this request */
	JanusErrorUnknownRequest = 453

	// JanusErrorInvalidJson /*! \brief The payload is not a valid JSON message */
	JanusErrorInvalidJson = 454

	// JanusErrorInvalidJsonObject /*! \brief The object is not a valid JSON object as expected */
	JanusErrorInvalidJsonObject = 455

	// JanusErrorMissingMandatoryElement /*! \brief A mandatory element is missing in the message */
	JanusErrorMissingMandatoryElement = 456

	// JanusErrorInvalidRequestPath /*! \brief The request cannot be handled for this webserver path  */
	JanusErrorInvalidRequestPath = 457

	// JanusErrorSessionNotFound /*! \brief The session the request refers to doesn't exist */
	JanusErrorSessionNotFound = 458

	// JanusErrorHandleNotFound /*! \brief The handle the request refers to doesn't exist */
	JanusErrorHandleNotFound = 459

	// JanusErrorPluginNotFound /*! \brief The plugin the request wants to talk to doesn't exist */
	JanusErrorPluginNotFound = 460

	// JanusErrorPluginAttach /*! \brief An error occurring when trying to attach to a plugin and create a handle  */
	JanusErrorPluginAttach = 461

	// JanusErrorPluginMessage /*! \brief An error occurring when trying to send a message/request to the plugin */
	JanusErrorPluginMessage = 462

	// JanusErrorPluginDetach /*! \brief An error occurring when trying to detach from a plugin and destroy the related handle  */
	JanusErrorPluginDetach = 463

	// JanusErrorJsepUnknownType /*! \brief The Janus core doesn't support this SDP type */
	JanusErrorJsepUnknownType = 464

	// JanusErrorJsepInvalidSdp /*! \brief The Session Description provided by the peer is invalid */
	JanusErrorJsepInvalidSdp = 465

	// JanusErrorTrickleInvalidStream /*! \brief The stream a trickle candidate for does not exist or is invalid */
	JanusErrorTrickleInvalidStream = 466

	// JanusErrorInvalidElementType /*! \brief A JSON element is of the wrong type (e.g., an integer instead of a string) */
	JanusErrorInvalidElementType = 467

	// JanusErrorSessionConflic /*! \brief The ID provided to create a new session is already in use */
	JanusErrorSessionConflic = 468

	// JanusErrorUnexpectedAnswer /*! \brief We got an ANSWER to an OFFER we never made */
	JanusErrorUnexpectedAnswer = 469

	// JanusErrorTokenNotFound /*! \brief The auth token the request refers to doesn't exist */
	JanusErrorTokenNotFound = 470

	// JanusErrorWebrtcState /*! \brief The current request cannot be handled because of not compatible WebRTC state */
	JanusErrorWebrtcState = 471

	// JanusErrorNotAcceptingSessions /*! \brief The server is currently configured not to accept new sessions */
	JanusErrorNotAcceptingSessions = 472
)
