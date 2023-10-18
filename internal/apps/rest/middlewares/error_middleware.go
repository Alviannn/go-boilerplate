package middlewares

import (
	"go-boilerplate/pkg/responses"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func CustomErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		customErr := responses.NewError().WithSourceError(err)

		log.Error().Err(customErr).Msg("Unhandled error")

		responses.New().
			WithError(customErr).
			Send(c)
	}
}
