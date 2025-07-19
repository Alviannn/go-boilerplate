//go:build !prod

package docs

import (
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"

	"github.com/labstack/echo/v4"
	echo_swagger "github.com/swaggo/echo-swagger"
)

func RegisterSwagger(app *echo.Echo) {
	if configs.Default().Environment != constants.EnvProduction {
		app.GET("/rest-swagger/*", echo_swagger.WrapHandler)
	}
}
