package logger

import (
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger() error {
	os.Mkdir("logs", os.ModePerm)

	zerolog.ErrorMarshalFunc = errorHandler
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var (
		consoleWriter io.Writer = os.Stdout
		config                  = configs.Default()
	)

	// When running in local environment (or basically not in production)
	// we'll enable debugging and pretty print logging.
	if config.Environment != constants.EnvProduction {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		consoleWriter = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: zerolog.TimeFieldFormat,
		}
	}

	rotateFileWriter, err := NewRotateFileWriter("./logs/{date}.log")
	if err != nil {
		return err
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, rotateFileWriter)
	log.Logger = log.Output(multi)
	return nil
}

func errorHandler(err error) any {
	customError, ok := err.(*customerror.Error)
	if !ok {
		return err
	}

	return customError.ToJSON()
}
