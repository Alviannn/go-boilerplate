package main

import (
	"context"
	"fmt"
	"go-boilerplate/internal/apps/rest/middlewares"
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/domains"
	domains_interfaces "go-boilerplate/internal/domains/interfaces"
	"go-boilerplate/pkg/customvalidator"
	"go-boilerplate/pkg/databases"
	"go-boilerplate/pkg/dependencies"

	"github.com/defval/di"
	"github.com/labstack/echo/v4"
	echo_middlewares "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	_ "go-boilerplate/internal/apps/rest/docs"
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

func StartServer(container *di.Container, validator *customvalidator.Validator) (err error) {
	if err = configs.Load(validator); err != nil {
		return
	}

	var gormDB *gorm.DB
	// Force DB to load and test the connection.
	if err = container.Resolve(&gormDB); err != nil {
		return
	}
	if err = databases.MigrateMySQL(); err != nil {
		return
	}

	app := echo.New()
	config := configs.Default()

	app.Use(echo_middlewares.Secure())
	app.Use(echo_middlewares.CORSWithConfig(echo_middlewares.CORSConfig{
		AllowOrigins: config.CORSAllowedOrigins,
	}))
	app.Pre(echo_middlewares.RemoveTrailingSlash())
	app.Use(echo_middlewares.RequestIDWithConfig(echo_middlewares.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, rid string) {
			req := c.Request()
			ctx := req.Context()

			newReq := req.WithContext(context.WithValue(ctx, constants.RequestIDKey, rid))
			c.SetRequest(newReq)
		},
	}))
	app.Use(middlewares.Log)

	// Override error handler middleware
	if err = RegisterRouters(app, container); err != nil {
		return
	}

	if config.Environment != constants.EnvProduction {
		app.GET("/rest-swagger/*", echoSwagger.WrapHandler)
	}

	app.HTTPErrorHandler = middlewares.CustomErrorHandler()
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
	container, err := dependencies.New(
		domains.Modules,
	)
	if err != nil {
		return
	}

	if err := container.Invoke(StartServer); err != nil {
		log.Fatal().Err(err).Send()
	}
}
