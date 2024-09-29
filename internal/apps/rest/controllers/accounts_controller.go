package controllers_rest

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type accounts struct {
	Service services.AccountsService
}

func NewAccounts(service services.AccountsService) *accounts {
	return &accounts{
		Service: service,
	}
}

func (ctl *accounts) SetupRouter(echo *echo.Echo) {
	router := echo.Group("accounts")

	router.GET("/:id", ctl.GetByID)
	router.GET("", ctl.GetAll)
	router.POST("", ctl.Register)
}

// GetByID gets account detail by the account ID
//
//	@Summary		Gets account details
//	@Description	Gets account details by the account ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account's ID"
//	@Success		200	{object}	models_mysql.Account
//	@Failure		400	{object}	customerror.ErrorJSON
//	@Failure		500	{object}	customerror.ErrorJSON
//	@Router			/accounts/{id} [get]
func (ctl *accounts) GetByID(c echo.Context) (err error) {
	var (
		params dtos.AccountGetRequest
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

	account, err := ctl.Service.GetByID(ctx, params)
	res = response.NewBuilder().
		WithSuccessCode(http.StatusOK).
		WithData(account).
		WithError(err).
		Build()

	return c.JSON(res.StatusCode, res.Data)
}

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
//	@Success		200			{object}	[]models_mysql.Account
//	@Failure		400			{object}	customerror.ErrorJSON
//	@Failure		500			{object}	customerror.ErrorJSON
//	@Router			/accounts [get]
func (ctl *accounts) GetAll(c echo.Context) (err error) {
	var (
		params dtos.AccountGetAllRequest
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

	data, err := ctl.Service.GetAll(ctx, params)
	res = response.NewBuilder().
		WithSuccessCode(http.StatusOK).
		WithData(data).
		WithError(err).
		Build()

	return c.JSON(res.StatusCode, res.Data)
}

// Register creates or registers a new account
//
//	@Summary		Register a new account
//	@Description	Creates or registers a new account
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			newAccount	body		dtos.AccountRegisterRequest	true	"New account details"
//	@Success		200			{object}	any
//	@Failure		400			{object}	customerror.ErrorJSON
//	@Failure		500			{object}	customerror.ErrorJSON
//	@Router			/accounts [post]
func (ctl *accounts) Register(c echo.Context) (err error) {
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

	err = ctl.Service.Register(ctx, params)
	res = response.NewBuilder().
		WithSuccessCode(http.StatusOK).
		WithError(err).
		Build()

	return c.JSON(res.StatusCode, res.Data)
}
