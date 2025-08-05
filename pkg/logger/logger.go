package logger

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type SetupWithConfigParam struct {
	ConsoleWriter io.Writer
}

func SetupWithConfig(param SetupWithConfigParam) error {
	var (
		writerList = []io.Writer{}
		cfg        = configs.Default()
	)

	err := os.Mkdir(cfg.LogsDir, os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	zerolog.ErrorMarshalFunc = errorMarshaller
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// When running in local environment (or basically not in production)
	// we'll enable debugging and pretty print logging.
	if cfg.Environment != constants.EnvProduction {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

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

	writerList = append(writerList, NewRotateFileWriter(func() string {
		fileName := fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly))
		return path.Join(cfg.LogsDir, fileName)
	}))

	log.Logger = log.Hook(&requestIDHook{})
	log.Logger = log.Output(zerolog.MultiLevelWriter(writerList...))
	return nil
}

func Setup() error {
	return SetupWithConfig(SetupWithConfigParam{
		ConsoleWriter: os.Stdout,
	})
}

func errorMarshaller(err error) any {
	customError, ok := err.(*customerror.Error)
	if !ok {
		return err
	}
	return customError.ToJSON()
}
