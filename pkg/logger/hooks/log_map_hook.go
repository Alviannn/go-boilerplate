package logger_hooks

import (
	"go-boilerplate/internal/constants"
	logger_types "go-boilerplate/pkg/logger/types"

	"github.com/rs/zerolog"
)

type LogMapHook struct{}

// Run implements the zerolog.Hook interface.
func (h *LogMapHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	var (
		ctx   = e.GetCtx()
		value = ctx.Value(constants.CtxKeyLogMap)
	)
	if value == nil {
		return
	}
	if logMap, ok := value.(logger_types.LogMapFunc); ok {
		logMap(e)
	}
}
