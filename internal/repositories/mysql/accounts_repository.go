package repositories_mysql

import (
	"context"
	"fmt"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"

	"gorm.io/gorm"
)

type accounts struct {
	DB *gorm.DB
}

func NewAccounts(db *gorm.DB) AccountsMySQLRepository {
	return &accounts{}
}

func (r *accounts) GetByID(ctx context.Context, accountID int64) (account models_mysql.Account, err error) {
	query := getDB(ctx, r.DB).
		WithContext(ctx).
		Limit(1).
		Where("id = ?", accountID)

	if err = query.Find(&account).Error; err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get account.")
		return
	}
	if !account.IsExist() {
		err = customerror.New().
			WithCode(http.StatusNotFound).
			WithMessage(fmt.Sprintf("Cannot find account with ID %d", accountID))
		return
	}
	return
}

func (r *accounts) GetAll(ctx context.Context, params dtos.AccountGetAllRequest) (accounts []models_mysql.Account, err error) {
	query := getDB(ctx, r.DB).WithContext(ctx)

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
		query = query.Limit(int(params.Limit))
	}
	if params.Offset != 0 {
		query = query.Offset(params.Offset)
	}

	if err = query.Find(&accounts).Error; err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get all accounts.")
	}
	return
}

func (r *accounts) Register(ctx context.Context, params dtos.AccountRegisterRequest) (err error) {
	newAccount := models_mysql.Account{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	query := getDB(ctx, r.DB).
		WithContext(ctx).
		Create(&newAccount)

	if err = query.Error; err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to register new account.")
	}
	return
}

func (r *accounts) ExistByEmail(ctx context.Context, email string) (exist bool, err error) {
	var account models_mysql.Account

	query := getDB(ctx, r.DB).
		WithContext(ctx).
		Select("id").
		Where("email = ?", email).
		Limit(1)

	if err = query.Find(&account).Error; err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to check account existence.")
		return
	}

	exist = account.IsExist()
	return
}
