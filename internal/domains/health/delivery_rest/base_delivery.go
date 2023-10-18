package health_delivery_rest

import (
	"github.com/labstack/echo/v4"
)

type restDeliveryImpl struct{}

func NewRestDelivery() *restDeliveryImpl {
	return &restDeliveryImpl{}
}

func (d *restDeliveryImpl) SetupRouter(app *echo.Echo) {
	group := app.Group("/health")

	group.GET("", d.GetHealth)
}
