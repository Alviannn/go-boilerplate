package repositories_mysql

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

type Tx interface {
	Transaction(ctx context.Context, fc func(newCtx context.Context) error) (err error)
}

type Accounts interface {
	GetByID(ctx context.Context, accountID int64) (account models_mysql.Account, err error)
	GetAll(ctx context.Context, params dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error)
	Register(ctx context.Context, params dtos.AccountRegisterReq) error
	ExistByEmail(ctx context.Context, email string) (exist bool, err error)
}
