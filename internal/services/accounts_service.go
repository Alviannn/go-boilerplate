package services

import (
	"context"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	repositories_mysql "go-boilerplate/internal/repositories/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type accounts struct {
	MySQLRepo repositories_mysql.Accounts
}

func NewAccounts(mysqlRepo repositories_mysql.Accounts) AccountsService {
	return &accounts{
		MySQLRepo: mysqlRepo,
	}
}

func (s *accounts) GetByID(ctx context.Context, params dtos.AccountGetReq) (account models_mysql.Account, err error) {
	account, err = s.MySQLRepo.GetByID(ctx, params.ID)
	return
}

func (s *accounts) GetAll(ctx context.Context, params dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error) {
	accounts, err = s.MySQLRepo.GetAll(ctx, params)
	return
}

func (s *accounts) Register(ctx context.Context, params dtos.AccountRegisterReq) (err error) {
	isExist, err := s.MySQLRepo.ExistByEmail(ctx, params.Email)
	if err != nil {
		return
	}
	if isExist {
		err = customerror.New().
			WithCode(http.StatusConflict).
			WithMessage("Account with this email already exists.")
		return
	}

	params.Password, err = s.hashPassword(params.Password)
	if err != nil {
		return
	}

	err = s.MySQLRepo.Register(ctx, params)
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
