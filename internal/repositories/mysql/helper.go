package repositories_mysql

import (
	"context"
	"go-boilerplate/internal/constants"

	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type helper struct {
	DB *gorm.DB `do:""`
}

func NewHelper(i do.Injector) (Helper, error) {
	return do.InvokeStruct[*helper](i)
}

func (r *helper) GetDB(ctx context.Context) (mysqlDB *gorm.DB) {
	mysqlDB = r.DB
	value := ctx.Value(constants.CtxKeyGormTransaction)

	if value == nil {
		return
	}

	txMysqlDB, ok := value.(*gorm.DB)
	if !ok {
		return
	}

	mysqlDB = txMysqlDB
	return
}

func (r *helper) Transaction(ctx context.Context, fc func(newCtx context.Context) error) error {
	db := r.GetDB(ctx)

	return db.Transaction(func(gormTx *gorm.DB) error {
		newCtx := context.WithValue(ctx, constants.CtxKeyGormTransaction, gormTx)
		return fc(newCtx)
	})
}
