package registeruser

import (
	"go-boilerplate/internal/dtos"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return Handler{Service: service}
}

func (h Handler) Handle(c echo.Context) (err error) {
	var params dtos.RegisterUserReq
	if err = c.Bind(&params); err != nil {
		return
	}

	return h.Service.RegisterUser(params)
}
