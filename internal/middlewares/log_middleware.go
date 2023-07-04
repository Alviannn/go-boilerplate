package middlewares

import (
	"fmt"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/helpers"
	"go-boilerplate/pkg/logger"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type LogHandler struct {
	rotateFileWriter *logger.RotateFileWriter
}

func NewLogHandler() (logHandler *LogHandler, err error) {
	requestLogDir := path.Join("logs", "requests")
	if err = os.MkdirAll(requestLogDir, os.ModePerm); err != nil {
		return
	}

	writer, err := logger.NewRotateFileWriter(path.Join(requestLogDir, "requests-{date}.log"))
	logHandler = &LogHandler{
		rotateFileWriter: writer,
	}
	return
}

func (h *LogHandler) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		startTime := time.Now()

		var clonedBody any
		if err := helpers.Echo.CloneBindedBody(c, &clonedBody); err != nil {
			clonedBody = nil
		}

		clonedRequest, err := helpers.Http.CloneRequest(c.Request())
		if err != nil {
			return
		}

		if err := next(c); err != nil {
			c.Error(err)
		}
		if err = h.dumpRequestToLog(clonedRequest, startTime); err != nil {
			return
		}

		elapsedTime := time.Since(startTime)
		err = h.logRequest(c, clonedBody, elapsedTime)
		return
	}
}

func (h *LogHandler) dumpRequestToLog(req *http.Request, startTime time.Time) (err error) {
	formattedTime := fmt.Sprintf("[%s]\n", startTime.Format(time.RFC3339))
	if _, err = h.rotateFileWriter.Write([]byte(formattedTime)); err != nil {
		return
	}

	if err = req.Write(h.rotateFileWriter); err != nil {
		return
	}

	_, err = h.rotateFileWriter.Write([]byte("\n\n\n"))
	return
}

func (h *LogHandler) logRequest(c echo.Context, bodyResult any, elapsedTime time.Duration) (err error) {
	req := c.Request()
	ctx := req.Context()

	reqHeader := c.Request().Header
	requestUri := string(req.RequestURI)
	httpMethod := string(req.Method)
	requestId := fmt.Sprint(ctx.Value(constants.RequestIDKey))

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
