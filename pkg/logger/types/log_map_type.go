package logger_types

import (
	"github.com/rs/zerolog"
)

type LogMapFunc func(*zerolog.Event)

func LogMapFromMap(m map[string]any) LogMapFunc {
	return func(e *zerolog.Event) {
		e.Fields(m)
	}
}

func LogMapFromSlice(s []any) LogMapFunc {
	return func(e *zerolog.Event) {
		e.Fields(s)
	}
}
