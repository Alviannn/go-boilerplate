package accounts_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	mysql_models "go-boilerplate/internal/models/mysql"
)

type MySQLRepository interface {
	GetByID(ctx context.Context, accountID int64) (account mysql_models.Account, err error)
	GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []mysql_models.Account, err error)
	Register(ctx context.Context, params dtos.RegisterAccountReq) error
	IsExistByEmail(ctx context.Context, email string) bool
}
