package main

import (
	"fmt"
	controllers_rest "go-boilerplate/internal/apps/rest/controllers"
	"go-boilerplate/internal/apps/rest/docs"
	"go-boilerplate/internal/apps/rest/middlewares"
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/databases"
	"go-boilerplate/pkg/dependencies"
	"go-boilerplate/pkg/helpers"
	"time"

	"github.com/defval/di"
	"github.com/labstack/echo/v4"
	echo_middlewares "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	_ "go-boilerplate/internal/apps/rest/docs"
)

func registerRouters(echo *echo.Echo, container *di.Container) (err error) {
	var restDeliveries []controllers_rest.Controller
	if err = container.Resolve(&restDeliveries); err != nil {
		return
	}

	for _, rest := range restDeliveries {
		rest.SetupRouter(echo)
	}
	return
}

func StartServer(container *di.Container) (err error) {
	// Force DB to load and test the connection.
	var gormDB *gorm.DB
	if err = container.Resolve(&gormDB); err != nil {
		return
	}
	if err = databases.MigrateMySQL(); err != nil {
		return
	}

	app := echo.New()
	config := configs.Default()

	app.Use(echo_middlewares.RecoverWithConfig(echo_middlewares.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			err = customerror.New().
				WithPanic(true).
				WithSourceError(err).
				WithMessage("PANIC: Unhandled error")
			return err
		},
	}))

	app.Pre(echo_middlewares.RemoveTrailingSlash())

	app.Use(echo_middlewares.Secure())
	app.Use(echo_middlewares.CORSWithConfig(echo_middlewares.CORSConfig{
		AllowOrigins: config.CORSAllowedOrigins,
	}))
	app.Use(echo_middlewares.RequestIDWithConfig(echo_middlewares.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, rid string) {
			helpers.EchoAddContextValue(c, constants.CtxKeyRequestID, rid)
		},
	}))
	app.Use(middlewares.NewLog().Handle)
	app.Use(middlewares.NewTimeout(30 * time.Second))

	app.HTTPErrorHandler = middlewares.CustomErrorHandler()
	app.JSONSerializer = middlewares.NewErrorGuardJSONSerializer(app)

	// Override error handler middleware
	if err = registerRouters(app, container); err != nil {
		return
	}

	docs.RegisterSwagger(app)

	err = app.Start(fmt.Sprintf(":%d", config.Port))
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

	container, err := dependencies.New(
		repositories.Module(),
		services.Module(),
		controllers_rest.Module(),
	)
	if err != nil {
		return
	}

	if err := container.Invoke(StartServer); err != nil {
		log.Fatal().Err(err).Send()
	}
}
