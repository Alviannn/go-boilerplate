package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger() error {
	os.Mkdir("logs", os.ModePerm)

	zerolog.ErrorMarshalFunc = HandleCustomError
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var consoleWriter io.Writer = os.Stdout

	// When running in local environment (or basically not in production)
	// we'll enable debugging and pretty print logging.
	if os.Getenv("ENVIRONMENT") != "production" {
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
