package middlewares

import (
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/response"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func CustomErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed { // Prevent sending double response
			return
		}

		var customErr *customerror.Error
		if tmpErr, ok := err.(*customerror.Error); ok && tmpErr.IsPanic {
			customErr = tmpErr
		} else {
			customErr = customerror.New().WithSourceError(err)
		}

		log.Error().Err(customErr).Msg(customErr.Message)

		res := response.NewBuilder().
			WithError(customErr).
			Build()

		c.JSON(res.StatusCode, res.Data)
	}
}
