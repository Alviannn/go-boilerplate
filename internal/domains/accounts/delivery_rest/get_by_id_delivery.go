package accounts_delivery_rest

import (
	"net/http"

	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"

	"github.com/labstack/echo/v4"
)

// GetByID gets account detail by the account ID
//
//	@Summary		Gets account details
//	@Description	Gets account details by the account ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account's ID"
//	@Success		200	{object}	models.Account
//	@Failure		400	{object}	responses.ErrorResponse
//	@Failure		500	{object}	responses.ErrorResponse
//	@Router			/accounts/{id} [get]
func (d *restDeliveryImpl) GetByID(c echo.Context) (err error) {
	var params dtos.GetAccountReq
	ctx := c.Request().Context()

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")

		return responses.New().WithError(err).Send(c)
	}

	account, err := d.Service.GetByID(ctx, params)
	return responses.New().
		WithData(account).
		WithError(err).
		WithSuccessCode(http.StatusOK).
		Send(c)
}
