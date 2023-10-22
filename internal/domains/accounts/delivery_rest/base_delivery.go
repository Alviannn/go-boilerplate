package accounts_delivery_rest

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"github.com/labstack/echo/v4"
)

type deliveryImpl struct {
	Service accounts_interfaces.Service
}

func New(service accounts_interfaces.Service) *deliveryImpl {
	return &deliveryImpl{
		Service: service,
	}
}

func (d *deliveryImpl) SetupRouter(echo *echo.Echo) {
	router := echo.Group("accounts")

	router.GET("/:id", d.GetByID)
	router.GET("", d.GetAll)
	router.POST("", d.Register)
}
