package accounts_delivery_rest

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"github.com/labstack/echo/v4"
)

type restDeliveryImpl struct {
	Service accounts_interfaces.Service
}

func NewRestDelivery(service accounts_interfaces.Service) *restDeliveryImpl {
	return &restDeliveryImpl{
		Service: service,
	}
}

func (d *restDeliveryImpl) SetupRouter(echo *echo.Echo) {
	router := echo.Group("accounts")

	router.GET("/:id", d.GetByID)
	router.GET("", d.GetAll)
	router.POST("", d.Register)
}
