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
//	@Failure		400			{object}	responses.ErrorResponse
//	@Failure		500			{object}	responses.ErrorResponse
//	@Router			/accounts [post]
func (d *restDeliveryImpl) Register(c echo.Context) (err error) {
	var (
		params dtos.RegisterAccountReq
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
