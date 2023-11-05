package accounts_delivery_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register creates or registers a new account
//
//	@Summary		Register a new account
//	@Description	Creates or registers a new account
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			newAccount	body		dtos.RegisterAccountReq	true	"New account details"
//	@Success		200			{object}	any
//	@Failure		400			{object}	customerror.ErrorJSON
//	@Failure		500			{object}	customerror.ErrorJSON
//	@Router			/accounts [post]
func (d *deliveryImpl) Register(c echo.Context) (err error) {
	var (
		params dtos.AccountRegisterRequest
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

	err = d.Service.Register(ctx, params)
	res = response.NewBuilder().
		WithSuccessCode(http.StatusOK).
		WithError(err).
		Build()

	return c.JSON(res.StatusCode, res.Data)
}
