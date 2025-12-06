package helpers

import (
	"context"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"
	"net/http"

	"github.com/labstack/echo/v4"
)

func EchoDefaultBind(c echo.Context, param any) (err error) {
	if err := c.Bind(param); err != nil {
		err = customerror.New().
			WithContext(c.Request().Context()).
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage(constants.ErrFailedBind)
	}
	return
}

func EchoAddContextValue(c echo.Context, key constants.ContextKey, value any) {
	req := c.Request()
	ctx := req.Context()

	newCtx := context.WithValue(ctx, key, value)
	newReq := req.WithContext(newCtx)

	c.SetRequest(newReq)
}
