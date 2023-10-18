package health_delivery_rest

import (
	"go-boilerplate/pkg/responses"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (d *restDeliveryImpl) GetHealth(c echo.Context) (err error) {
	currentTime := time.Now().Format(time.DateTime)
	return responses.New().
		WithData(currentTime).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
