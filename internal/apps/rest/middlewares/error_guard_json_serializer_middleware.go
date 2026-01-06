package middlewares

import (
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/helpers"

	"github.com/labstack/echo/v4"
)

// ErrorGuardJSONSerializer is a serializer that intercepts the `customerror.Error`
// before being written to `ResponseWriter` through the use of `c.JSON` method.
//
// What is does is:
//  1. Prevent the error stack trace (or `Stack`) from being shown in the
//     HTTP response for security reasons.
//  2. Allows the error stack trace (or `Stack`) to be shown in the logs.
type ErrorGuardJSONSerializer struct {
	Parent echo.JSONSerializer
}

func NewErrorGuardJSONSerializer(app *echo.Echo) echo.JSONSerializer {
	return &ErrorGuardJSONSerializer{
		Parent: app.JSONSerializer,
	}
}

func (t *ErrorGuardJSONSerializer) Serialize(c echo.Context, v any, indent string) error {
	if errResp, ok := v.(customerror.ErrorJSON); ok {
		// Store original error in `context` so it can log
		// with stack trace in the logger middleware.
		helpers.EchoAddContextValue(c, constants.CtxKeyHTTPTraceableError, errResp)

		// Prevent the error stack trace from being shown in HTTP response.
		if configs.Default().Environment == constants.EnvProduction {
			errResp.Stack = nil
			v = errResp
		}
	}

	return t.Parent.Serialize(c, v, indent)
}

func (t *ErrorGuardJSONSerializer) Deserialize(c echo.Context, v any) error {
	return t.Parent.Deserialize(c, v)
}
