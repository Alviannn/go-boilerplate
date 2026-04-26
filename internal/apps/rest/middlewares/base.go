package middlewares

import (
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/helpers"
	"time"

	"github.com/labstack/echo/v4"
	echo_middlewares "github.com/labstack/echo/v4/middleware"
)

func Use(app *echo.Echo) {
	app.Use(echo_middlewares.RecoverWithConfig(echo_middlewares.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			err = customerror.New().
				WithPanic(true).
				WithSourceError(err).
				WithMessage("PANIC: Unhandled error")
			return err
		},
	}))

	app.Pre(echo_middlewares.RemoveTrailingSlash())

	app.Use(echo_middlewares.Secure())
	app.Use(echo_middlewares.CORSWithConfig(echo_middlewares.CORSConfig{
		AllowOrigins: configs.Default().CORSAllowedOrigins,
	}))
	app.Use(echo_middlewares.RequestIDWithConfig(echo_middlewares.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, rid string) {
			helpers.EchoAddContextValue(c, constants.CtxKeyRequestID, rid)
		},
	}))
	app.Use(Log())
	app.Use(echo_middlewares.ContextTimeout(30 * time.Second))

	app.HTTPErrorHandler = CustomErrorHandler()
	app.JSONSerializer = NewErrorGuardJSONSerializer(app)
}
