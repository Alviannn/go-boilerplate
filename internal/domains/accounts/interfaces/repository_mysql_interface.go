package accounts_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

type RepositoryMySQL interface {
	Transaction(deps dtos.TxDependencies) RepositoryMySQL

	GetByID(ctx context.Context, accountID int64) (account models_mysql.Account, err error)
	GetAll(ctx context.Context, params dtos.AccountGetAllRequest) (accounts []models_mysql.Account, err error)
	Register(ctx context.Context, params dtos.AccountRegisterRequest) error
	ExistByEmail(ctx context.Context, email string) (exist bool, err error)
}
