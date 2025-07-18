package middlewares

import (
	"context"
	"go-boilerplate/pkg/customerror"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func NewTimeout(timeout time.Duration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				ctx     = c.Request().Context()
				errChan = make(chan error, 1)
			)

			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			c.SetRequest(c.Request().WithContext(ctx))

			// Use goroutine to prevent blocking the main thread
			go func() {
				errChan <- next(c)
			}()

			select {
			case err := <-errChan:
				c.Error(err)
			case <-ctx.Done():
				err := customerror.New().
					WithSourceError(ctx.Err()).
					WithCode(http.StatusServiceUnavailable).
					WithMessage("Request Timeout")
				c.Error(err)
			}

			return nil
		}
	}
}
