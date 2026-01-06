package middlewares

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

var (
	MaskedValue = strings.Repeat("*", 10)

	FieldToMask = []string{
		"password",
		"token",
		"access_token",
		"refresh_token",
	}
)

type logMiddleware struct{}

func Log() echo.MiddlewareFunc {
	logMiddleware := &logMiddleware{}
	return logMiddleware.Handle
}

func (m logMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	startTime := time.Now()

	bodyDumpConfig := echo_middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			var (
				uri         = c.Request().RequestURI
				elapsedTime = time.Since(startTime)
			)

			if strings.Contains(uri, "swagger") {
				return
			}

			m.writeLogRequest(c, reqBody)
			m.writeLogResponse(c, resBody, elapsedTime)
		},
	}

	return echo_middleware.BodyDumpWithConfig(bodyDumpConfig)(next)
}

func (m logMiddleware) writeLogRequest(c echo.Context, buf []byte) {
	var (
		body any

		req    = c.Request()
		ctx    = req.Context()
		header = req.Header.Clone()
	)

	if buf != nil {
		if err := json.Unmarshal(buf, &body); err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("failed to unmarshal request body")
		}
	}

	m.maskHeaders(header)

	log.Info().
		Ctx(ctx).
		Str("method", req.Method).
		Str("uri", req.RequestURI).
		Any("body", m.maskJSONBody(body)).
		Any("headers", header).
		Msg("HTTP Request")
}

func (m logMiddleware) writeLogResponse(c echo.Context, buf []byte, elapsedTime time.Duration) {
	var (
		body any

		req    = c.Request()
		ctx    = req.Context()
		res    = c.Response()
		header = res.Header().Clone()
	)

	logEvent := log.Info()
	if res.Status >= http.StatusBadRequest {
		logEvent = log.Error()
	}

	if traceableVal := ctx.Value(constants.CtxKeyHTTPTraceableError); traceableVal != nil {
		var err error

		traceableErr, ok := traceableVal.(customerror.ErrorJSON)
		if !ok {
			err = fmt.Errorf("invalid traceable error type")
			log.Error().Ctx(ctx).Err(err).Msg(err.Error())
		}

		buf, err = json.Marshal(traceableErr)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("failed to marshal traceable error")
		}
	}

	if buf != nil {
		if err := json.Unmarshal(buf, &body); err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("failed to unmarshal response body")
			body = nil
		}
	}

	m.maskHeaders(header)

	logEvent.
		Ctx(ctx).
		Str("method", req.Method).
		Str("uri", req.RequestURI).
		Str("elapsed_time", fmt.Sprintf("%dms", elapsedTime.Milliseconds())).
		Int("status_code", res.Status).
		Any("body", m.maskJSONBody(body)).
		Any("headers", header).
		Msg("HTTP Response")
}

func (m logMiddleware) maskJSONBody(body any) any {
	switch typedBody := body.(type) {
	case []any:
		for i, item := range typedBody {
			typedBody[i] = m.maskJSONBody(item)
		}

		return typedBody
	case map[string]any:
		for key, value := range typedBody {
			isShouldMask := slices.Contains(FieldToMask, key)
			if _, ok := value.(string); ok && isShouldMask {
				typedBody[key] = MaskedValue
				continue
			}

			typedBody[key] = m.maskJSONBody(value)
		}

		return typedBody
	default:
		return body
	}
}

func (m logMiddleware) maskHeaders(header http.Header) {
	m.maskCookies(header)
	m.maskAuthorization(header)
}

func (m logMiddleware) maskCookies(header http.Header) {
	var (
		matchRegexFmt = `%s=[a-zA-Z0-9._-]+`
		headerKeyList = []string{
			echo.HeaderSetCookie,
			echo.HeaderCookie,
		}
	)

	for _, key := range headerKeyList {
		for idx, value := range header[key] {
			for _, field := range FieldToMask {
				regex, err := regexp.Compile(fmt.Sprintf(matchRegexFmt, field))
				if err != nil {
					continue
				}

				header[key][idx] = regex.ReplaceAllString(
					value,
					fmt.Sprintf(
						"%s=%s",
						field,
						MaskedValue,
					),
				)
			}
		}
	}
}

func (m logMiddleware) maskAuthorization(header http.Header) {
	authorization := header.Get(echo.HeaderAuthorization)
	if !strings.HasPrefix(authorization, "Bearer ") {
		return
	}

	header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", MaskedValue))
}
