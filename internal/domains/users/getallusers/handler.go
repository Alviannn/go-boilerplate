package getallusers

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
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
	var params dtos.GetAllUsersReq

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")
		return
	}

	data, err := h.Service.GetAllUsers(params)
	return responses.New().
		WithData(data).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
