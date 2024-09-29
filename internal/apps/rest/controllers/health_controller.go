package controllers_rest

import (
	"go-boilerplate/pkg/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type health struct{}

func NewHealth() *health {
	return &health{}
}

func (ctl *health) SetupRouter(app *echo.Echo) {
	group := app.Group("/health")

	group.GET("", ctl.Get)
}

func (ctl *health) Get(c echo.Context) (err error) {
	currentTime := time.Now().Format(time.DateTime)
	res := response.NewBuilder().
		WithData(currentTime).
		WithSuccessCode(http.StatusOK).
		Build()

	return c.JSON(res.StatusCode, res.Data)
}
