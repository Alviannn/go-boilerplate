package users_delivery_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterUser creates or registers a new user
//
//	@Summary		Register a new user
//	@Description	Creates or registers a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			newUser	body		dtos.RegisterUserReq	true	"New user details"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	responses.ErrorResponse
//	@Failure		500		{object}	responses.ErrorResponse
//	@Router			/users [post]
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
