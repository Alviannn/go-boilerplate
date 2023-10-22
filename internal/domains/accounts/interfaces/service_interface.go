package accounts_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	mysql_models "go-boilerplate/internal/models/mysql"
)

type Service interface {
	GetByID(ctx context.Context, params dtos.GetAccountReq) (account mysql_models.Account, err error)
	GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []mysql_models.Account, err error)
	Register(ctx context.Context, params dtos.RegisterAccountReq) (err error)
}
