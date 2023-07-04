package main

import (
	"context"
	"fmt"
	"go-boilerplate/internal/constants"
	domains_interfaces "go-boilerplate/internal/domains/interfaces"
	"go-boilerplate/internal/middlewares"
	"go-boilerplate/pkg/dependencies"
	"go-boilerplate/pkg/logger"
	"os"

	"github.com/goava/di"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echo_middlewares "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func RegisterRouters(echo *echo.Echo, container *di.Container) (err error) {
	var restDeliveries []domains_interfaces.BaseRestDelivery
	if err = container.Resolve(&restDeliveries); err != nil {
		return
	}

	for _, rest := range restDeliveries {
		rest.SetupRouter(echo)
	}
	return
}

func StartServer() (err error) {
	if err = godotenv.Load(); err != nil {
		return
	}
	if err = logger.SetupLogger(); err != nil {
		return
	}

	// Set logging for dependency registery and resolving.
	di.SetTracer(&logger.DITracer{})

	container, err := dependencies.New()
	if err != nil {
		return
	}

	err = container.Invoke(func(app *echo.Echo) (err error) {
		// Force DB to load and test the connection.
		var gorm *gorm.DB
		if err = container.Resolve(&gorm); err != nil {
			return
		}

		app.Use(echo_middlewares.RemoveTrailingSlash())
		app.Use(echo_middlewares.RequestIDWithConfig(echo_middlewares.RequestIDConfig{
			RequestIDHandler: func(c echo.Context, rid string) {
				req := c.Request()
				ctx := req.Context()

				ctx = context.WithValue(ctx, constants.RequestIDKey, rid)
				c.SetRequest(req.WithContext(ctx))
			},
		}))

		// Override error handler middleware
		if err = RegisterRouters(app, container); err != nil {
			return
		}

		app.HTTPErrorHandler = middlewares.ErrorHandler()
		err = app.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
		return
	})
	return
}

func main() {
	if err := StartServer(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
