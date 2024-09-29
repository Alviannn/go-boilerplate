package middlewares

import (
	"context"
	"go-boilerplate/internal/configs"
	"go-boilerplate/internal/constants"
	"go-boilerplate/pkg/customerror"

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

func (t *ErrorGuardJSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	if customErrorJSON, ok := i.(customerror.ErrorJSON); ok {
		req := c.Request()
		ctx := req.Context()

		// Store original error in `context` so it can log
		// with stack trace in the logger middleware.
		newCtx := context.WithValue(ctx, constants.HTTPTraceableError, customErrorJSON)
		newReq := req.WithContext(newCtx)

		c.SetRequest(newReq)

		// Prevent the error stack trace from being shown in HTTP response.
		if configs.Default().Environment == constants.EnvProduction {
			customErrorJSON.Stack = nil
			i = customErrorJSON
		}
	}

	return t.Parent.Serialize(c, i, indent)
}

func (t *ErrorGuardJSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	return t.Parent.Deserialize(c, i)
}
