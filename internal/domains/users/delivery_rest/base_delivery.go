package users_delivery_rest

import (
	users_interfaces "go-boilerplate/internal/domains/users/interfaces"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type RestDeliveryImpl struct {
	di.Inject

	Service users_interfaces.Service
}

func NewRestDelivery(p RestDeliveryImpl) users_interfaces.RestDelivery {
	return &p
}

func (d *RestDeliveryImpl) SetupRouter(echo *echo.Echo) {
	router := echo.Group("users")

	router.GET("/:id", d.GetUser)
	router.GET("", d.GetAllUsers)
	router.POST("", d.RegisterUser)
}
