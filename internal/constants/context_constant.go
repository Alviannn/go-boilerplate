package constants

type ContextKey string

const (
	RequestIDKey       = ContextKey("requestid")
	GormTransactionKey = ContextKey("gorm_transaction")
)
