package accounts_delivery_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/response"
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
//	@Failure		400			{object}	customerror.ErrorJSON
//	@Failure		500			{object}	customerror.ErrorJSON
//	@Router			/accounts [get]
func (d *restDeliveryImpl) GetAll(c echo.Context) (err error) {
	var (
		params dtos.GetAllAccountsReq
		res    response.Response

		ctx = c.Request().Context()
	)

	if err = c.Bind(&params); err != nil {
		err = customerror.New().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")

		res = response.NewBuilder().WithError(err).Build()
		return c.JSON(res.StatusCode, res.Data)
	}

	data, err := d.Service.GetAll(ctx, params)
	res = response.NewBuilder().
		WithSuccessCode(http.StatusOK).
		WithData(data).
		WithError(err).
		Build()

	return c.JSON(res.StatusCode, res.Data)
}
