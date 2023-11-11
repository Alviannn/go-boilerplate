package middlewares

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/helpers"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	startTime := time.Now()

	bodyDumpMiddleware := echoMiddleware.BodyDumpWithConfig(echoMiddleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqRawBody, resRawBody []byte) {
			uri := c.Request().RequestURI
			if strings.Contains(uri, "swagger") {
				return
			}

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

	requestID := fmt.Sprint(req.Context().Value(constants.RequestIDKey))
	body = maskJSONBody(body)

	header := c.Request().Header.Clone()
	maskHeaders(header)

	log.Info().
		Str("request_id", requestID).
		Str("method", req.Method).
		Str("uri", req.RequestURI).
		Any("body", body).
		Any("headers", header).
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

	header := res.Header().Clone()
	maskHeaders(header)

	logEvent.
		Str("request_id", requestID).
		Str("method", req.Method).
		Str("uri", req.RequestURI).
		Str("elapsed_time", fmt.Sprintf("%dms", elapsedTime.Milliseconds())).
		Int("status_code", res.Status).
		Any("body", body).
		Any("headers", header).
		Msg("HTTP Response")
}

func maskJSONBody(body any) any {
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
				typedBody[key] = constants.MaskedValue
				continue
			}

			typedBody[key] = maskJSONBody(value)
		}

		return typedBody
	default:
		return body
	}
}

func maskHeaders(header http.Header) {
	maskCookies(header)
	maskAuthorization(header)
}

func maskCookies(header http.Header) {
	cookieFields := []string{
		"refresh_token",
	}

	for mapKey, mapValue := range header {
		for arrayIndex, arrayValue := range mapValue {
			for _, field := range cookieFields {
				regex, err := regexp.Compile(fmt.Sprintf(`%s=[a-zA-Z0-9._-]+`, field))
				if err != nil {
					continue
				}

				header[mapKey][arrayIndex] = regex.ReplaceAllString(arrayValue, fmt.Sprintf("%s=%s", field, constants.MaskedValue))
			}
		}
	}
}

func maskAuthorization(header http.Header) {
	authorization := header.Get("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		return
	}

	header.Set("Authorization", fmt.Sprintf("Bearer %s", constants.MaskedValue))
}

func unmarshalAnyOrNil(jsonValue []byte) (value any) {
	if err := json.Unmarshal(jsonValue, &value); err != nil {
		value = nil
	}
	return
}
