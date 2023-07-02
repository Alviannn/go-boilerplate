package users_delivery_rest

import (
	"net/http"

	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"

	"github.com/labstack/echo/v4"
)

func (d *RestDeliveryImpl) GetUser(c echo.Context) (err error) {
	var params dtos.GetUserReq
	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	user, err := d.Service.GetUser(params)
	return responses.New().
		WithData(user).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
