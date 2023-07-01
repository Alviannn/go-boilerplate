package logger

import "github.com/rs/zerolog/log"

type DITracer struct {
}

func (*DITracer) Trace(format string, args ...interface{}) {
	log.Debug().Msgf(format, args...)
}
