package health_delivery_rest

import (
	"github.com/labstack/echo/v4"
)

type deliveryImpl struct{}

func New() *deliveryImpl {
	return &deliveryImpl{}
}

func (d *deliveryImpl) SetupRouter(app *echo.Echo) {
	group := app.Group("/health")

	group.GET("", d.Get)
}
