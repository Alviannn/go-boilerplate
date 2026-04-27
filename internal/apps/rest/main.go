package main

import (
	"errors"
	"fmt"
	controllers_rest "go-boilerplate/internal/apps/rest/controllers"
	"go-boilerplate/internal/apps/rest/docs"
	"go-boilerplate/internal/apps/rest/middlewares"
	"go-boilerplate/internal/configs"
	"go-boilerplate/pkg/databases"
	"go-boilerplate/pkg/logger"
	"io"
	"os"
	"path"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"gorm.io/gorm"

	_ "go-boilerplate/internal/apps/rest/docs"
)

func setupLogger() (err error) {
	cfg := configs.Default()

	err = os.MkdirAll(cfg.LogsDir, os.ModePerm)
	if errors.Is(err, os.ErrExist) {
		err = nil
	}
	if err != nil {
		return
	}

	fileWriter := logger.NewRotateFileWriter(func() string {
		fileName := fmt.Sprintf("%s.log", time.Now().Format(time.DateOnly))
		return path.Join(cfg.LogsDir, fileName)
	})

	logger.Setup(logger.SetupParam{
		ConsoleWriter: os.Stdout,
		ExtraWriters:  []io.Writer{fileWriter},
	})
	return
}

func startServer(injector *do.RootScope) (err error) {
	// Force DB to load and test the connection.
	_, err = do.Invoke[*gorm.DB](injector)
	if err != nil {
		return
	}

	if err = databases.MigrateMySQL(); err != nil {
		return
	}

	app := echo.New()

	middlewares.Use(app)

	// Register all controllers
	if err = controllers_rest.Register(injector, app); err != nil {
		return
	}

	docs.RegisterSwagger(app)

	err = app.Start(fmt.Sprintf(":%d", configs.Default().Port))
	return
}

// main starts the REST API server
//
//	@title		API documentation
//	@version	1.0
//	@schemes	http https
//	@host		localhost:5000
func main() {
	if err := configs.Load(); err != nil {
		panic(err)
	}

	if err := setupLogger(); err != nil {
		panic(err)
	}

	injector := NewDI()

	if err := startServer(injector); err != nil {
		panic(err)
	}
}
