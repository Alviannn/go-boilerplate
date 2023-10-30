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
	body = maskJSONBody(body)

	log.Info().
		Str("request_id", requestID).
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
	body = maskJSONBody(body)

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

func maskJSONBody(body any) any {
	maskedValue := "************"
	fieldsToMask := []string{
		"password",
		"token",
		"access_token",
		"refresh_token",
	}

	switch typedBody := body.(type) {
	case []any:
		for i, item := range typedBody {
			typedBody[i] = maskJSONBody(item)
		}

		return typedBody
	case map[string]any:
		for key, value := range typedBody {
			isShouldMask := helpers.Slice[string]().IsIn(fieldsToMask, key)
			if _, ok := value.(string); ok && isShouldMask {
				typedBody[key] = maskedValue
				continue
			}

			typedBody[key] = maskJSONBody(value)
		}

		return typedBody
	default:
		return body
	}
}

func unmarshalAnyOrNil(jsonValue []byte) (value any) {
	if err := json.Unmarshal(jsonValue, &value); err != nil {
		value = nil
	}
	return
}
