package accounts_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

type RepositoryMySQL interface {
	GetByID(ctx context.Context, accountID int64) (account models_mysql.Account, err error)
	GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []models_mysql.Account, err error)
	Register(ctx context.Context, params dtos.RegisterAccountReq) error
	IsExistByEmail(ctx context.Context, email string) (exist bool, err error)
}
