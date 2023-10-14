package middlewares

import (
	"fmt"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/helpers"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		startTime := time.Now()

		var clonedBody any
		if err := helpers.Echo.CloneBindedBody(c, &clonedBody); err != nil {
			clonedBody = nil
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		elapsedTime := time.Since(startTime)
		err = logRequest(c, clonedBody, elapsedTime)
		return
	}
}

func logRequest(c echo.Context, bodyResult any, elapsedTime time.Duration) (err error) {
	req := c.Request()

	reqHeader := c.Request().Header
	requestUri := string(req.RequestURI)
	httpMethod := string(req.Method)
	requestId := fmt.Sprint(c.Get(constants.RequestIDKey))

	log.Info().
		Str("requestId", requestId).
		Str("method", httpMethod).
		Str("uri", requestUri).
		Str("elapsed", fmt.Sprintf("%dms", elapsedTime.Milliseconds())).
		Any("body", bodyResult).
		Any("header", reqHeader).
		Msg("HTTP Request")
	return
}
