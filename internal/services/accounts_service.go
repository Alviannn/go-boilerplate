package services

import (
	"context"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	repositories_mysql "go-boilerplate/internal/repositories/mysql"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/customvalidator"
	"net/http"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type accounts struct {
	MySQLRepo repositories_mysql.Accounts
	Validator *customvalidator.Validator
}

func NewAccounts(
	mysqlRepo repositories_mysql.Accounts,
	validator *customvalidator.Validator,
) Accounts {
	return &accounts{
		MySQLRepo: mysqlRepo,
		Validator: validator,
	}
}

func (s *accounts) GetByID(ctx context.Context, param dtos.AccountGetReq) (account models_mysql.Account, err error) {
	if err = s.Validator.Validate(&param); err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusBadRequest).
			WithMessage(constants.ErrFailedValidation)
		return
	}

	account, err = s.MySQLRepo.Get(ctx, models_mysql.AccountGetParam{
		ID: param.ID,
	})
	if err != nil {
		return
	}
	return
}

func (s *accounts) GetAll(ctx context.Context, param dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error) {
	accounts, err = s.MySQLRepo.GetAll(ctx, param)
	if err != nil {
		return
	}
	return
}

func (s *accounts) Register(ctx context.Context, param dtos.AccountRegisterReq) (err error) {
	if err = s.Validator.Validate(&param); err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusBadRequest).
			WithMessage(constants.ErrFailedValidation)
		return
	}

	_, err = s.MySQLRepo.Get(ctx, models_mysql.AccountGetParam{
		Email: param.Email,
	})
	if err == nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusConflict).
			WithMessage("Account with this email already exists")
		return
	}
	if customerror.EqualCode(err, http.StatusNotFound) {
		err = nil
	}
	if err != nil {
		return
	}

	hashedPassword, err := s.hashPassword(ctx, param.Password)
	if err != nil {
		return
	}

	account := models_mysql.Account{
		Username: param.Username,
		FullName: param.FullName,
		Email:    param.Email,
		Password: hashedPassword,
	}

	err = s.MySQLRepo.Create(ctx, &account)
	if err != nil {
		return
	}
	return
}

func (s *accounts) hashPassword(ctx context.Context, password string) (hashed string, err error) {
	rawHashed, err := bcrypt.GenerateFromPassword([]byte(password), constants.DefaultHashCost)
	if err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Failed to hash password")

		log.Error().Ctx(ctx).Err(err).Send()
		return
	}

	hashed = string(rawHashed)
	return
}
