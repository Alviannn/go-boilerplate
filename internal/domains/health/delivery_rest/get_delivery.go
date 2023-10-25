package health_delivery_rest

import (
	"go-boilerplate/pkg/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (d *deliveryImpl) Get(c echo.Context) (err error) {
	currentTime := time.Now().Format(time.DateTime)
	res := response.NewBuilder().
		WithData(currentTime).
		WithSuccessCode(http.StatusOK).
		Build()

	return c.JSON(res.StatusCode, res.Data)
}
