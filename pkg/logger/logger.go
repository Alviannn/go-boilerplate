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

func Setup() error {
	var (
		config        = configs.Default()
		consoleWriter = io.Writer(os.Stdout)
	)

	if err := os.Mkdir(config.LogsDir, os.ModePerm); err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	zerolog.ErrorMarshalFunc = errorMarshaller
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// When running in local environment (or basically not in production)
	// we'll enable debugging and pretty print logging.
	if config.Environment != constants.EnvProduction {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		consoleWriter = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: zerolog.TimeFieldFormat,
		}
	}

	rotateFileWriter, err := NewRotateFileWriter(func() string {
		fileName := fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly))
		return path.Join(config.LogsDir, fileName)
	})
	if err != nil {
		return err
	}

	log.Logger = log.Output(zerolog.MultiLevelWriter(
		consoleWriter,
		rotateFileWriter,
	))
	return nil
}

func errorMarshaller(err error) any {
	customError, ok := err.(*customerror.Error)
	if !ok {
		return err
	}
	return customError.ToJSON()
}
