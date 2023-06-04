package middlewares

import (
	"fmt"
	"go-boilerplate/pkg/responses"

	"github.com/labstack/echo/v4"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		resp := responses.New()

		if echoError, ok := err.(*echo.HTTPError); ok {
			newError := responses.NewError().
				WithCode(echoError.Code).
				WithError(echoError).
				WithMessage(fmt.Sprint(echoError.Message))

			resp.WithError(newError)
		} else {
			resp.WithError(err)
		}

		resp.Send(c)
	}
}
