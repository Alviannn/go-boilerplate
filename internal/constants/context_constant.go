package constants

type ContextKey string

const (
	CtxKeyRequestID          = ContextKey("requestid")
	CtxKeyGormTransaction    = ContextKey("gorm_transaction")
	CtxKeyHTTPTraceableError = ContextKey("http_traceable_error")
)
