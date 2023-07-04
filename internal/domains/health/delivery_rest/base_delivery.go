package health_delivery_rest

import (
	health_interfaces "go-boilerplate/internal/domains/health/interfaces"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type RestDeliveryImpl struct {
	di.Inject
}

func NewRestDelivery(p RestDeliveryImpl) health_interfaces.RestDelivery {
	return &p
}

func (d *RestDeliveryImpl) SetupRouter(app *echo.Echo) {
	group := app.Group("/health")

	group.GET("", d.GetHealth)
}
