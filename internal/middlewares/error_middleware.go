package middlewares

import (
	"go-boilerplate/pkg/responses"

	"github.com/labstack/echo/v4"
)

func CustomErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		customErr := responses.NewError().
			WithSourceError(err)

		responses.New().
			WithError(customErr).
			Send(c)
	}
}
