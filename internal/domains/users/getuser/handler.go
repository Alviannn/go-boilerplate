package getuser

import (
	"go-boilerplate/internal/dtos"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return Handler{Service: service}
}

func (h Handler) Handle(c echo.Context) (err error) {
	var params dtos.GetUserReq
	if err = c.Bind(&params); err != nil {
		return
	}

	user, err := h.Service.GetUser(params)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, user)
}
