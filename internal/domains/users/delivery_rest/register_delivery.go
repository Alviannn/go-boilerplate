package users_delivery_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d *RestDeliveryImpl) RegisterUser(c echo.Context) (err error) {
	var params dtos.RegisterUserReq
	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")

		return
	}

	return d.Service.RegisterUser(params)
}
