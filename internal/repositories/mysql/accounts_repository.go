package repositories_mysql

import (
	"context"
	"errors"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type accounts struct {
	RepositoryMysql
}

func NewAccounts(repository RepositoryMysql) Accounts {
	return &accounts{
		RepositoryMysql: repository,
	}
}

func (r *accounts) Get(ctx context.Context, params models_mysql.AccountGetParam) (account models_mysql.Account, err error) {
	query := r.GetDB(ctx).
		WithContext(ctx).
		Model(&models_mysql.Account{})

	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	if params.Username != "" {
		query = query.Where("username = ?", params.Username)
	}
	if params.Email != "" {
		query = query.Where("email = ?", params.Email)
	}

	err = query.First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = customerror.New().
			WithContext(ctx).
			WithSourceError(err).
			WithCode(http.StatusNotFound).
			WithMessage("Account not found")
		return
	}
	if err != nil {
		err = customerror.New().
			WithContext(ctx).
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get account")

		log.Error().Ctx(ctx).Err(err).Send()
		return
	}
	return
}

func (r *accounts) GetAll(ctx context.Context, params dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error) {
	query := r.GetDB(ctx).
		WithContext(ctx).
		Model(&models_mysql.Account{})

	if params.Email != "" {
		query = query.Where("email = ?", params.Email)
	}
	if params.FullName != "" {
		query = query.Where("full_name = ?", params.FullName)
	}
	if params.Username != "" {
		query = query.Where("username = ?", params.Username)
	}

	if params.Limit != 0 {
		query = query.Limit(params.Limit)
		if params.Offset != 0 {
			query = query.Offset(params.Offset)
		}
	}

	if err = query.Find(&accounts).Error; err != nil {
		err = customerror.New().
			WithContext(ctx).
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get all accounts")

		log.Error().Ctx(ctx).Err(err).Send()
		return
	}
	return
}

func (r *accounts) Create(ctx context.Context, account *models_mysql.Account) (err error) {
	query := r.GetDB(ctx).
		WithContext(ctx).
		Model(&models_mysql.Account{}).
		Create(account)

	if err = query.Error; err != nil {
		err = customerror.New().
			WithContext(ctx).
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to create new account")

		log.Error().Ctx(ctx).Err(err).Send()
		return
	}
	return
}
