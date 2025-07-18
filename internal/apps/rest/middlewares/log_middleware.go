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
)

type Log struct{}

func NewLog() *Log {
	return &Log{}
}

func (m *Log) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	startTime := time.Now()

	bodyDumpConfig := echo_middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqRawBody, resRawBody []byte) {
			uri := c.Request().RequestURI
			if strings.Contains(uri, "swagger") {
				return
			}

			reqBody := m.unmarshalAnyOrNil(reqRawBody)
			m.writeLogRequest(c, reqBody)

			elapsedTime := time.Since(startTime)

			resBody := m.unmarshalAnyOrNil(resRawBody)
			if errorRespBody := m.getTraceableErrorAsAnyOrNil(c); errorRespBody != nil {
				resBody = errorRespBody
			}

			m.writeLogResponse(c, resBody, elapsedTime)
		},
	}

	return echo_middleware.BodyDumpWithConfig(bodyDumpConfig)(next)
}

func (m *Log) writeLogRequest(c echo.Context, body any) {
	var (
		req    = c.Request()
		header = req.Header.Clone()
	)

	m.maskHeaders(header)

	log.Info().
		Ctx(req.Context()).
		Str("method", req.Method).
		Str("uri", req.RequestURI).
		Any("body", m.maskJSONBody(body)).
		Any("headers", header).
		Msg("HTTP Request")
}

func (m *Log) writeLogResponse(c echo.Context, body any, elapsedTime time.Duration) {
	var (
		req    = c.Request()
		res    = c.Response()
		header = res.Header().Clone()

		isErrorResponse = (res.Status >= http.StatusBadRequest)
	)

	logEvent := log.Info()
	if isErrorResponse {
		logEvent = log.Error()
	}

	m.maskHeaders(header)

	logEvent.
		Ctx(req.Context()).
		Str("method", req.Method).
		Str("uri", req.RequestURI).
		Str("elapsed_time", fmt.Sprintf("%dms", elapsedTime.Milliseconds())).
		Int("status_code", res.Status).
		Any("body", m.maskJSONBody(body)).
		Any("headers", header).
		Msg("HTTP Response")
}

func (m *Log) maskJSONBody(body any) any {
	fieldToMaskList := []string{
		"password",
		"token",
		"access_token",
		"refresh_token",
	}

	switch typedBody := body.(type) {
	case []any:
		for i, item := range typedBody {
			typedBody[i] = m.maskJSONBody(item)
		}

		return typedBody
	case map[string]any:
		for key, value := range typedBody {
			isShouldMask := slices.Contains(fieldToMaskList, key)
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

func (m *Log) maskHeaders(header http.Header) {
	m.maskCookies(header)
	m.maskAuthorization(header)
}

func (m *Log) maskCookies(header http.Header) {
	var (
		matchRegexFmt = `%s=[a-zA-Z0-9._-]+`
		headerKeyList = []string{
			echo.HeaderSetCookie,
			echo.HeaderCookie,
		}
		fieldList = []string{
			"refresh_token",
		}
	)

	for _, key := range headerKeyList {
		for idx, value := range header[key] {
			for _, field := range fieldList {
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

func (m *Log) maskAuthorization(header http.Header) {
	authorization := header.Get(echo.HeaderAuthorization)
	if !strings.HasPrefix(authorization, "Bearer ") {
		return
	}

	header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", MaskedValue))
}

func (m *Log) unmarshalAnyOrNil(jsonValue []byte) (value any) {
	if err := json.Unmarshal(jsonValue, &value); err != nil {
		value = nil
	}
	return
}

func (m *Log) getTraceableErrorAsAnyOrNil(c echo.Context) any {
	ctx := c.Request().Context()
	value := ctx.Value(constants.CtxKeyHTTPTraceableError)
	if value == nil {
		return nil
	}

	traceableError, ok := value.(customerror.ErrorJSON)
	if !ok {
		return nil
	}

	marshaledError, err := json.Marshal(&traceableError)
	if err != nil {
		return nil
	}

	return m.unmarshalAnyOrNil(marshaledError)
}
