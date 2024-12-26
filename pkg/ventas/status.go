package ventas

type Status string

const (
	StatusContinue                      Status = "100 CONTINUE"
	StatusSwitchingProtocols            Status = "101 SWITCHING_PROTOCOLS"
	StatusProcessing                    Status = "102 PROCESSING"
	StatusEarlyHints                    Status = "103 EARLY_HINTS"
	StatusCheckpoint                    Status = "103 CHECKPOINT"
	StatusOK                            Status = "200 OK"
	StatusCreated                       Status = "201 CREATED"
	StatusAccepted                      Status = "202 ACCEPTED"
	StatusNonAuthoritativeInformation   Status = "203 NON_AUTHORITATIVE_INFORMATION"
	StatusNoContent                     Status = "204 NO_CONTENT"
	StatusResetContent                  Status = "205 RESET_CONTENT"
	StatusPartialContent                Status = "206 PARTIAL_CONTENT"
	StatusMultiStatus                   Status = "207 MULTI_STATUS"
	StatusAlreadyReported               Status = "208 ALREADY_REPORTED"
	StatusIMUsed                        Status = "226 IM_USED"
	StatusMultipleChoices               Status = "300 MULTIPLE_CHOICES"
	StatusMovedPermanently              Status = "301 MOVED_PERMANENTLY"
	StatusFound                         Status = "302 FOUND"
	StatusMovedTemporarily              Status = "302 MOVED_TEMPORARILY"
	StatusSeeOther                      Status = "303 SEE_OTHER"
	StatusNotModified                   Status = "304 NOT_MODIFIED"
	StatusUseProxy                      Status = "305 USE_PROXY"
	StatusTemporaryRedirect             Status = "307 TEMPORARY_REDIRECT"
	StatusPermanentRedirect             Status = "308 PERMANENT_REDIRECT"
	StatusBadRequest                    Status = "400 BAD_REQUEST"
	StatusUnauthorized                  Status = "401 UNAUTHORIZED"
	StatusPaymentRequired               Status = "402 PAYMENT_REQUIRED"
	StatusForbidden                     Status = "403 FORBIDDEN"
	StatusNotFound                      Status = "404 NOT_FOUND"
	StatusMethodNotAllowed              Status = "405 METHOD_NOT_ALLOWED"
	StatusNotAcceptable                 Status = "406 NOT_ACCEPTABLE"
	StatusProxyAuthenticationRequired   Status = "407 PROXY_AUTHENTICATION_REQUIRED"
	StatusRequestTimeout                Status = "408 REQUEST_TIMEOUT"
	StatusConflict                      Status = "409 CONFLICT"
	StatusGone                          Status = "410 GONE"
	StatusLengthRequired                Status = "411 LENGTH_REQUIRED"
	StatusPreconditionFailed            Status = "412 PRECONDITION_FAILED"
	StatusPayloadTooLarge               Status = "413 PAYLOAD_TOO_LARGE"
	StatusRequestEntityTooLarge         Status = "413 REQUEST_ENTITY_TOO_LARGE"
	StatusURITooLong                    Status = "414 URI_TOO_LONG"
	StatusRequestURITooLong             Status = "414 REQUEST_URI_TOO_LONG"
	StatusUnsupportedMediaType          Status = "415 UNSUPPORTED_MEDIA_TYPE"
	StatusRequestedRangeNotSatisfiable  Status = "416 REQUESTED_RANGE_NOT_SATISFIABLE"
	StatusExpectationFailed             Status = "417 EXPECTATION_FAILED"
	StatusIAmATeapot                    Status = "418 I_AM_A_TEAPOT"
	StatusInsufficientSpaceOnResource   Status = "419 INSUFFICIENT_SPACE_ON_RESOURCE"
	StatusMethodFailure                 Status = "420 METHOD_FAILURE"
	StatusDestinationLocked             Status = "421 DESTINATION_LOCKED"
	StatusUnprocessableEntity           Status = "422 UNPROCESSABLE_ENTITY"
	StatusLocked                        Status = "423 LOCKED"
	StatusFailedDependency              Status = "424 FAILED_DEPENDENCY"
	StatusTooEarly                      Status = "425 TOO_EARLY"
	StatusUpgradeRequired               Status = "426 UPGRADE_REQUIRED"
	StatusPreconditionRequired          Status = "428 PRECONDITION_REQUIRED"
	StatusTooManyRequests               Status = "429 TOO_MANY_REQUESTS"
	StatusRequestHeaderFieldsTooLarge   Status = "431 REQUEST_HEADER_FIELDS_TOO_LARGE"
	StatusUnavailableForLegalReasons    Status = "451 UNAVAILABLE_FOR_LEGAL_REASONS"
	StatusInternalServerError           Status = "500 INTERNAL_SERVER_ERROR"
	StatusNotImplemented                Status = "501 NOT_IMPLEMENTED"
	StatusBadGateway                    Status = "502 BAD_GATEWAY"
	StatusServiceUnavailable            Status = "503 SERVICE_UNAVAILABLE"
	StatusGatewayTimeout                Status = "504 GATEWAY_TIMEOUT"
	StatusHTTPVersionNotSupported       Status = "505 HTTP_VERSION_NOT_SUPPORTED"
	StatusVariantAlsoNegotiates         Status = "506 VARIANT_ALSO_NEGOTIATES"
	StatusInsufficientStorage           Status = "507 INSUFFICIENT_STORAGE"
	StatusLoopDetected                  Status = "508 LOOP_DETECTED"
	StatusBandwidthLimitExceeded        Status = "509 BANDWIDTH_LIMIT_EXCEEDED"
	StatusNotExtended                   Status = "510 NOT_EXTENDED"
	StatusNetworkAuthenticationRequired Status = "511 NETWORK_AUTHENTICATION_REQUIRED"
)