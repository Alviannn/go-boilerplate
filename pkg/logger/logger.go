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
		logsDir       = "./logs"
		config        = configs.Default()
		consoleWriter = io.Writer(os.Stdout)
	)

	mkdirErr := os.Mkdir(logsDir, os.ModePerm).(*os.PathError)
	if mkdirErr != nil && !errors.Is(mkdirErr.Err, os.ErrExist) {
		return mkdirErr
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
		fileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
		return path.Join(logsDir, fileName)
	})
	if err != nil {
		return err
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, rotateFileWriter)
	log.Logger = log.Output(multi)
	return nil
}

func errorMarshaller(err error) any {
	customError, ok := err.(*customerror.Error)
	if !ok {
		return err
	}
	return customError.ToJSON()
}
