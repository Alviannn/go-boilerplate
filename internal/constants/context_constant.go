package constants

type ContextKey string

const (
	RequestIDKey       = ContextKey("requestid")
	GormTransactionKey = ContextKey("gorm_transaction")
	HTTPTraceableError = ContextKey("http_traceable_error")
)
