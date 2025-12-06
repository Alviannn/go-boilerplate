package hooks

import (
	"fmt"
	"go-boilerplate/internal/constants"

	"github.com/rs/zerolog"
)

type RequestIDHook struct{}

// Run implements the zerolog.Hook interface.
func (h *RequestIDHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	ctx := e.GetCtx()
	if requestID := ctx.Value(constants.CtxKeyRequestID); requestID != nil {
		e.Str("request_id", fmt.Sprint(requestID))
	}
}
