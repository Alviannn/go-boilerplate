package accounts_delivery_rest

import (
	"net/http"

	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/response"

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
	var (
		params dtos.GetAccountReq
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

	account, err := d.Service.GetByID(ctx, params)
	res = response.NewBuilder().
		WithSuccessCode(http.StatusOK).
		WithData(account).
		WithError(err).
		Build()

	return c.JSON(res.StatusCode, res.Data)
}
