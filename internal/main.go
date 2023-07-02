package main

import (
	"fmt"
	domains_interfaces "go-boilerplate/internal/domains/interfaces"
	"go-boilerplate/internal/middlewares"
	"go-boilerplate/pkg/dependencies"
	"go-boilerplate/pkg/logger"
	"os"

	"github.com/goava/di"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

func main() {
	if err := godotenv.Load(); err != nil {
		return
	}
	if err := logger.SetupLogger(); err != nil {
		log.Fatal().Err(err).Msg("Failed to setup logger configuration.")
	}

	// Set logging for dependency registery and resolving.
	di.SetTracer(&logger.DITracer{})

	container, err := dependencies.New()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = container.Invoke(func(echo *echo.Echo) (err error) {
		// Force DB to load and test the connection.
		var gorm *gorm.DB
		if err = container.Resolve(&gorm); err != nil {
			return
		}

		// Override error handler middleware
		if err = RegisterRouters(echo, container); err != nil {
			return
		}

		echo.HTTPErrorHandler = middlewares.ErrorHandler()
		err = echo.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
		return
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
