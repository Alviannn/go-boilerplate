package users_delivery_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllUsers gets user detail by the user ID
//
//	@Summary		Gets all users along with its details
//	@Description	Gets all users along with its details
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	false	"User's username"
//	@Param			email		query		string	false	"User's email address"
//	@Param			fullName	query		string	false	"User's full name"
//	@Param			limit		query		int		false	"Limit the amount data to show"
//	@Param			offset		query		int		false	"The data offset, or where it should start"
//	@Success		200			{object}	[]models.User
//	@Failure		400			{object}	responses.ErrorResponse
//	@Failure		500			{object}	responses.ErrorResponse
//	@Router			/users [get]
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
