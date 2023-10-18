package accounts_delivery_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
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
func (d *RestDeliveryImpl) Register(c echo.Context) (err error) {
	var params dtos.RegisterAccountReq
	ctx := c.Request().Context()

	if err = c.Bind(&params); err != nil {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithSourceError(err).
			WithMessage("Failed to bind parameters")

		return responses.New().WithError(err).Send(c)
	}

	err = d.Service.Register(ctx, params)
	return responses.New().
		WithError(err).
		Send(c)
}
