package middlewares

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/helpers"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	startTime := time.Now()

	bodyDumpMiddleware := echoMiddleware.BodyDumpWithConfig(echoMiddleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqRawBody, resRawBody []byte) {
			reqBody := unmarshalAnyOrNil(reqRawBody)
			writeLogRequest(c, reqBody)

			elapsedTime := time.Since(startTime)
			resBody := unmarshalAnyOrNil(resRawBody)
			writeLogResponse(c, resBody, elapsedTime)
		},
	})

	return bodyDumpMiddleware(next)
}

func writeLogRequest(c echo.Context, body any) {
	req := c.Request()

	reqHeader := c.Request().Header
	requestID := fmt.Sprint(req.Context().Value(constants.RequestIDKey))
	body = maskBody(body)

	log.Info().
		Str("requestId", requestID).
		Str("method", req.Method).
		Str("uri", req.RequestURI).
		Any("body", body).
		Any("headers", reqHeader).
		Msg("HTTP Request")
}

func writeLogResponse(c echo.Context, body any, elapsedTime time.Duration) {
	req := c.Request()
	ctx := req.Context()
	res := c.Response()

	requestID := fmt.Sprint(ctx.Value(constants.RequestIDKey))
	body = maskBody(body)

	logEvent := log.Info()
	isErrorResponse := (res.Status >= 400 && res.Status <= 500)

	if isErrorResponse {
		logEvent = log.Error()
	}

	logEvent.
		Str("request_id", requestID).
		Str("method", req.Method).
		Str("uri", req.RequestURI).
		Str("elapsed_time", fmt.Sprintf("%dms", elapsedTime.Milliseconds())).
		Int("status_code", res.Status).
		Any("body", body).
		Any("headers", res.Header()).
		Msg("HTTP Response")
}

func maskBody(body any) any {
	maskedValue := "************"
	fieldsToMask := []string{
		"password",
		"token",
		"access_token",
		"refresh_token",
	}

	if sliceBody, ok := body.([]any); ok {
		for i, item := range sliceBody {
			sliceBody[i] = maskBody(item)
		}
		return sliceBody
	}

	mapBody, ok := body.(map[string]any)
	if !ok {
		return body
	}

	for key, value := range mapBody {
		if innerMap, ok := mapBody[key].(map[string]any); ok {
			mapBody[key] = maskBody(innerMap)
		}

		isShouldMask := helpers.Slice[string]().IsIn(fieldsToMask, key)
		if _, ok := value.(string); ok && isShouldMask {
			mapBody[key] = maskedValue
		}
	}
	return mapBody
}

func unmarshalAnyOrNil(jsonValue []byte) (value any) {
	if err := json.Unmarshal(jsonValue, &value); err != nil {
		value = nil
	}
	return
}
