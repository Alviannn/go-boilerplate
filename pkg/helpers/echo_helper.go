package helpers

import (
	"go-boilerplate/pkg/customerror"
	"net/http"

	"github.com/labstack/echo/v4"
)

type echoHelper struct{}

func (echoHelper) DefaultBind(c echo.Context, param any) (err error) {
	if err := c.Bind(param); err != nil {
		err = customerror.New().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")
	}
	return
}
