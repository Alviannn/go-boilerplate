package accounts_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	postgres_models "go-boilerplate/internal/models/postgres"
)

type Repository interface {
	GetByID(ctx context.Context, accountID int64) (account postgres_models.Account, err error)
	GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []postgres_models.Account, err error)
	Register(ctx context.Context, params dtos.RegisterAccountReq) error
	IsExistByEmail(ctx context.Context, email string) bool
}
