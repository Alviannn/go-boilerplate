package services

import (
	"context"
	"crypto/sha256"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	repositories_mysql "go-boilerplate/internal/repositories/mysql"
	"go-boilerplate/pkg/customerror"
	"go-boilerplate/pkg/customvalidator"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/samber/do/v2"
	"golang.org/x/crypto/bcrypt"
)

type accounts struct {
	Validator *customvalidator.Validator `do:""`
	RepoMysql *repositories_mysql.Base   `do:""`
}

func NewAccounts(i do.Injector) (Accounts, error) {
	return do.InvokeStruct[*accounts](i)
}

func (s *accounts) GetByID(ctx context.Context, param dtos.AccountGetReq) (account models_mysql.Account, err error) {
	if err = s.Validator.Validate(&param); err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusBadRequest).
			WithMessage(constants.ErrFailedValidation)
		return
	}

	account, err = s.RepoMysql.Accounts.Get(ctx, models_mysql.AccountGetParam{
		ID: param.ID,
	})
	if err != nil {
		return
	}
	return
}

func (s *accounts) GetAll(ctx context.Context, param dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error) {
	accounts, err = s.RepoMysql.Accounts.GetAll(ctx, param)
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

	_, err = s.RepoMysql.Accounts.Get(ctx, models_mysql.AccountGetParam{
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

	err = s.RepoMysql.Helper.Transaction(ctx, func(newCtx context.Context) error {
		txErr := s.RepoMysql.Accounts.Create(newCtx, &account)
		if txErr != nil {
			return txErr
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}

func (s *accounts) hashPassword(ctx context.Context, password string) (hashed string, err error) {
	// pre-hash with sha256 before bcrypt for better security and allow longer passwords
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(password))

	rawHashed, err := bcrypt.GenerateFromPassword(sha256Hash.Sum(nil), constants.DefaultHashCost)
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
