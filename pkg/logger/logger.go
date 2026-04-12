package logger

import (
	"go-boilerplate/internal/configs"
	"go-boilerplate/pkg/logger/hooks"
	"io"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type SetupParam struct {
	ConsoleWriter io.Writer
	ExtraWriters  []io.Writer
}

func Setup(param SetupParam) {
	var (
		writerList  = []io.Writer{}
		globalLevel = zerolog.InfoLevel // defaults to info level
	)

	// When running in local environment (or basically not in production)
	// we'll enable debugging and pretty print logging.
	if !configs.Default().IsEnvProd() {
		globalLevel = zerolog.DebugLevel

		if param.ConsoleWriter != nil {
			param.ConsoleWriter = zerolog.ConsoleWriter{
				Out:        param.ConsoleWriter,
				TimeFormat: zerolog.TimeFieldFormat,
			}
		}
	}

	if param.ConsoleWriter != nil {
		writerList = append(writerList, param.ConsoleWriter)
	}
	if len(param.ExtraWriters) > 0 {
		writerList = append(writerList, param.ExtraWriters...)
	}

	zerolog.SetGlobalLevel(globalLevel)
	zerolog.ErrorMarshalFunc = hooks.ErrorMarshallerHook

	logWriter := zerolog.MultiLevelWriter(writerList...)

	newLogger := zerolog.New(logWriter).With().Timestamp().Logger()
	newLogger = newLogger.Hook(&hooks.RequestIDHook{})

	log.Logger = newLogger
}
