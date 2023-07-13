package users_delivery_rest

import (
	"net/http"

	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"

	"github.com/labstack/echo/v4"
)

// GetUser gets user detail by the user ID
//
//	@Summary		Gets user details
//	@Description	Gets user details by the user ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User's ID"
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	responses.ErrorResponse
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/users/{id} [get]
func (d *RestDeliveryImpl) GetUser(c echo.Context) (err error) {
	var params dtos.GetUserReq
	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")

		return responses.New().WithError(err).Send(c)
	}

	user, err := d.Service.GetUser(params)
	return responses.New().
		WithData(user).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
