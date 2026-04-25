package helpers

import (
	"context"
	"go-boilerplate/internal/constants"
	logger_types "go-boilerplate/pkg/logger/types"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func CtxAppendLogMap(ctx context.Context, logMap logger_types.LogMapFunc) (outCtx context.Context) {
	defer func() {
		outCtx = context.WithValue(ctx, constants.CtxKeyLogMap, logMap)
	}()

	value := ctx.Value(constants.CtxKeyLogMap)
	if value == nil {
		// No existing log map, just use the new one
		return
	}

	oldLogMap, ok := value.(logger_types.LogMapFunc)
	if !ok {
		log.Error().Ctx(ctx).Msg("Failed to append log map: expected LogMapFunc")
		return
	}

	logMapFromParam := logMap

	logMap = logger_types.LogMapFunc(func(e *zerolog.Event) {
		oldLogMap(e)
		logMapFromParam(e)
	})
	return
}
