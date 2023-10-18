package accounts_delivery_rest

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type RestDeliveryImpl struct {
	di.Inject

	Service accounts_interfaces.Service
}

func NewRestDelivery(p RestDeliveryImpl) *RestDeliveryImpl {
	return &p
}

func (d *RestDeliveryImpl) SetupRouter(echo *echo.Echo) {
	router := echo.Group("accounts")

	router.GET("/:id", d.GetByID)
	router.GET("", d.GetAll)
	router.POST("", d.Register)
}
