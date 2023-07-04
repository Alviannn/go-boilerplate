package helpers

import (
	"github.com/labstack/echo/v4"
)

type echoHelper struct{}

func (h *echoHelper) CloneBindedBody(c echo.Context, i any) error {
	clonedReq, err := Http.CloneRequest(c.Request())
	if err != nil {
		return err
	}

	clonedEchoCtx := c.Echo().NewContext(clonedReq, nil)
	return clonedEchoCtx.Bind(&i)
}
