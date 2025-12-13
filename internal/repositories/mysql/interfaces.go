package repositories_mysql

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"

	"gorm.io/gorm"
)

type RepositoryMysql interface {
	GetDB(ctx context.Context) (mysqlDB *gorm.DB)
}

type Tx interface {
	Transaction(ctx context.Context, fc func(newCtx context.Context) error) (err error)
}

type Accounts interface {
	Get(ctx context.Context, params models_mysql.AccountGetParam) (account models_mysql.Account, err error)
	GetAll(ctx context.Context, params dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error)
	Create(ctx context.Context, account *models_mysql.Account) (err error)
}
