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
			WithMessage("Failed to validate request")
		return
	}

	account, err = s.MySQLRepo.GetByID(ctx, param.ID)
	return
}

func (s *accounts) GetAll(ctx context.Context, param dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error) {
	accounts, err = s.MySQLRepo.GetAll(ctx, param)
	return
}

func (s *accounts) Register(ctx context.Context, param dtos.AccountRegisterReq) (err error) {
	if err = s.Validator.Validate(&param); err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Failed to validate request")
		return
	}

	isExist, err := s.MySQLRepo.ExistByEmail(ctx, param.Email)
	if err != nil {
		return
	}
	if isExist {
		err = customerror.New().
			WithCode(http.StatusConflict).
			WithMessage("Account with this email already exists.")
		return
	}

	param.Password, err = s.hashPassword(param.Password)
	if err != nil {
		return
	}

	err = s.MySQLRepo.Register(ctx, param)
	return
}

func (s *accounts) hashPassword(password string) (hashed string, err error) {
	rawHashed, err := bcrypt.GenerateFromPassword([]byte(password), constants.DefaultHashCost)
	if err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithMessage("Failed to hash password.").
			WithCode(http.StatusBadRequest)
		return
	}

	hashed = string(rawHashed)
	return
}
