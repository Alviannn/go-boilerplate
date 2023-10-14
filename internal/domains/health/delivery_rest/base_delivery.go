package health_delivery_rest

import (
	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type RestDeliveryImpl struct {
	di.Inject
}

func NewRestDelivery(p RestDeliveryImpl) *RestDeliveryImpl {
	return &p
}

func (d *RestDeliveryImpl) SetupRouter(app *echo.Echo) {
	group := app.Group("/health")

	group.GET("", d.GetHealth)
}
