package health_interfaces

import (
	domains_interfaces "go-boilerplate/internal/domains/interfaces"

	"github.com/labstack/echo/v4"
)

type RestDelivery interface {
	domains_interfaces.BaseRestDelivery

	GetHealth(c echo.Context) (err error)
}
