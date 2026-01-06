package logger

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/configs"
	"go-boilerplate/pkg/logger/hooks"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type SetupWithConfigParam struct {
	ConsoleWriter io.Writer
	IsDisableFile bool
}

func SetupWithConfig(param SetupWithConfigParam) error {
	var (
		writerList = []io.Writer{}
		cfg        = configs.Default()
	)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// When running in local environment (or basically not in production)
	// we'll enable debugging and pretty print logging.
	if !cfg.IsEnvProd() {
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

	if !param.IsDisableFile {
		err := os.Mkdir(cfg.LogsDir, os.ModePerm)
		if errors.Is(err, os.ErrExist) {
			err = nil
		}
		if err != nil {
			return err
		}

		writerList = append(writerList, NewRotateFileWriter(func() string {
			fileName := fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly))
			return path.Join(cfg.LogsDir, fileName)
		}))
	}

	applyHooksToLogger()

	log.Logger = log.Output(zerolog.MultiLevelWriter(writerList...))
	return nil
}

func applyHooksToLogger() {
	zerolog.ErrorMarshalFunc = hooks.ErrorMarshallerHook

	log.Logger = log.Hook(&hooks.RequestIDHook{})
}

func Setup() error {
	return SetupWithConfig(SetupWithConfigParam{
		ConsoleWriter: os.Stdout,
		IsDisableFile: false,
	})
}
