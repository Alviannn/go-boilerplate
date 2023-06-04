package getuser

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
	var params dtos.GetUserReq

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	user, err := h.Service.GetUser(params)
	return responses.New().
		WithData(user).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
