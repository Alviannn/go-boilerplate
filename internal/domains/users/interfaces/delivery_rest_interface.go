package users_interfaces

import (
	domains_interfaces "go-boilerplate/internal/domains/interfaces"

	"github.com/labstack/echo/v4"
)

type RestDelivery interface {
	domains_interfaces.BaseRestDelivery

	GetUser(c echo.Context) (err error)
	GetAllUsers(c echo.Context) (err error)
	RegisterUser(c echo.Context) (err error)
}
