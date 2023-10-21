package main

import (
	"context"
	"fmt"
	"go-boilerplate/internal/apps/rest/middlewares"
	"go-boilerplate/internal/constants"
	domains_interfaces "go-boilerplate/internal/domains/interfaces"
	"go-boilerplate/pkg/databases"
	"go-boilerplate/pkg/dependencies"
	"os"
	"strings"

	"github.com/defval/di"
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
		if err = databases.MigratePosgres(); err != nil {
			return
		}

		app.Use(echo_middlewares.Secure())
		app.Use(echo_middlewares.CORSWithConfig(echo_middlewares.CORSConfig{
			AllowOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		}))
		app.Use(echo_middlewares.RemoveTrailingSlash())
		app.Use(echo_middlewares.RequestIDWithConfig(echo_middlewares.RequestIDConfig{
			RequestIDHandler: func(c echo.Context, rid string) {
				req := c.Request()
				ctx := req.Context()

				newReq := req.WithContext(context.WithValue(ctx, constants.RequestIDKey, rid))
				c.SetRequest(newReq)
			},
		}))
		app.Use(middlewares.LogMiddleware)

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

// main starts the REST API server
//
//	@title		API documentation
//	@version	1.0
//	@schemes	http https
//	@host		localhost:5000
func main() {
	if err := StartServer(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
