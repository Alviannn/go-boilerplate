package accounts_delivery_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAll gets all accounts along with its details
//
//	@Summary		Gets all accounts along with its details
//	@Description	Gets all accounts along with its details
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	false	"Account's username"
//	@Param			email		query		string	false	"Account's email address"
//	@Param			fullName	query		string	false	"Account's full name"
//	@Param			limit		query		int		false	"Limit the amount data to show"
//	@Param			offset		query		int		false	"The data offset, or where it should start"
//	@Success		200			{object}	[]models.Account
//	@Failure		400			{object}	responses.ErrorResponse
//	@Failure		500			{object}	responses.ErrorResponse
//	@Router			/accounts [get]
func (d *restDeliveryImpl) GetAll(c echo.Context) (err error) {
	var params dtos.GetAllAccountsReq
	ctx := c.Request().Context()

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")

		return responses.New().WithError(err).Send(c)
	}

	data, err := d.Service.GetAll(ctx, params)
	return responses.New().
		WithData(data).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
