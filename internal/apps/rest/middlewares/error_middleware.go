package middlewares

import (
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/response"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func CustomErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		customErr := customerror.New().WithSourceError(err)
		log.Error().Err(customErr).Msg("Unhandled error")

		res := response.NewBuilder().WithError(customErr).Build()
		c.JSON(res.StatusCode, res.Data)
	}
}
