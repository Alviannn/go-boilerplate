package main

import (
	"context"
	"fmt"
	"go-boilerplate/internal/constants"
	domains_interfaces "go-boilerplate/internal/domains/interfaces"
	"go-boilerplate/internal/middlewares"
	"go-boilerplate/pkg/dependencies"
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

	container, err := dependencies.NewForStartup()
	if err != nil {
		return
	}

	err = container.Invoke(func(app *echo.Echo) (err error) {
		// Force DB to load and test the connection.
		var gorm *gorm.DB
		if err = container.Resolve(&gorm); err != nil {
			return
		}

		logMiddleware, err := middlewares.NewLogHandler()
		if err != nil {
			return
		}

		app.Use(echo_middlewares.RemoveTrailingSlash())
		app.Use(echo_middlewares.RequestIDWithConfig(echo_middlewares.RequestIDConfig{
			RequestIDHandler: func(c echo.Context, rid string) {
				req := c.Request()
				ctx := req.Context()

				ctx = context.WithValue(ctx, constants.RequestIDKey, rid)
				c.Set(constants.RequestIDKey, rid)
				c.SetRequest(req.WithContext(ctx))
			},
		}))
		app.Use(logMiddleware.Handler)

		// Override error handler middleware
		if err = RegisterRouters(app, container); err != nil {
			return
		}

		app.HTTPErrorHandler = middlewares.CustomErrorHandler()
		err = app.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
		return
	})
	return
}

//	@title		API documentation
//	@version	1.0
//	@schemes	http https
//	@host		localhost:5000
func main() {
	if err := StartServer(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
