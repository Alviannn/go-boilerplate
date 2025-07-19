//go:build prod

package docs

import "github.com/labstack/echo/v4"

func RegisterSwagger(app *echo.Echo) {
	// no-op for production
}
