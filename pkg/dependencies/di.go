package dependencies

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/configs"
	"go-boilerplate/pkg/customvalidator"
	"go-boilerplate/pkg/databases"
	"go-boilerplate/pkg/logger"
	"io"
	"os"
	"path"
	"time"

	"github.com/defval/di"
)

// New creates a new DI (dependency injection) container.
func New(extraDeps ...di.Option) (container *di.Container, err error) {
	fileWriter, err := prepareRotateFileWriter()
	if err != nil {
		return
	}

	logger.Setup(logger.SetupParam{
		ConsoleWriter: os.Stdout,
		ExtraWriters:  []io.Writer{fileWriter},
	})

	// Set logging for dependency registery and resolving.
	di.SetTracer(&logger.DITracer{})

	deps := []di.Option{
		di.Provide(customvalidator.New),
		di.Provide(databases.NewMySQLDB),
	}
	deps = append(deps, extraDeps...)

	return di.New(deps...)
}

func prepareRotateFileWriter() (fileWriter io.Writer, err error) {
	cfg := configs.Default()

	err = os.MkdirAll(cfg.LogsDir, os.ModePerm)
	if errors.Is(err, os.ErrExist) {
		err = nil
	}
	if err != nil {
		return
	}

	fileWriter = logger.NewRotateFileWriter(func() string {
		fileName := fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly))
		return path.Join(cfg.LogsDir, fileName)
	})
	return
}
