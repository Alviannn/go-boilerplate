package users_delivery_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d *RestDeliveryImpl) GetAllUsers(c echo.Context) (err error) {
	var params dtos.GetAllUsersReq
	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")
		return
	}

	data, err := d.Service.GetAllUsers(params)
	return responses.New().
		WithData(data).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
