package controllers_rest

import "github.com/labstack/echo/v4"

type Controller interface {
	SetupRouter(app *echo.Echo)
}
