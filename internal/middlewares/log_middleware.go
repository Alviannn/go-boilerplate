package middlewares

import (
	"fmt"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/helpers"
	"strings"
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
	requestId := fmt.Sprint(req.Context().Value(constants.RequestIDKey))
	bodyResult = maskBody(bodyResult)

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

func maskBody(body any) any {
	maskedValue := "************"

	if sliceBody, ok := body.([]any); ok {
		for i, item := range sliceBody {
			sliceBody[i] = maskBody(item)
		}
		return sliceBody
	}

	mapBody := body.(map[string]any)
	for key, value := range mapBody {
		if innerMap, ok := mapBody[key].(map[string]any); ok {
			mapBody[key] = maskBody(innerMap)
		}
		if _, ok := value.(string); strings.ToLower(key) == "password" && ok {
			mapBody[key] = maskedValue
		}
	}
	return mapBody
}
