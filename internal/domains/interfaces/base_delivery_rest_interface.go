package domains_interfaces

import "github.com/labstack/echo/v4"

type BaseRestDelivery interface {
	SetupRouter(app *echo.Echo)
}
