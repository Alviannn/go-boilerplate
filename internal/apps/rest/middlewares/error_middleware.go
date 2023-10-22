package middlewares

import (
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/response"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func CustomErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		var customError *customerror.Error
		if castableError, ok := err.(*customerror.Error); ok {
			customError = castableError
		} else {
			customError = customerror.New().WithSourceError(err)
		}

		log.Error().Err(customError).Msg("Unhandled error")

		res := response.NewBuilder().WithError(customError).Build()
		c.JSON(res.StatusCode, res.Data)
	}
}
