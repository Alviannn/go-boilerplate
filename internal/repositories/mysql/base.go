package repositories_mysql

import (
	"context"
	"go-boilerplate/internal/constants"

	"github.com/defval/di"
	"gorm.io/gorm"
)

func Module() di.Option {
	return di.Options(
		di.Provide(NewTx),
		di.Provide(NewAccounts),
	)
}

func getDB(ctx context.Context, db *gorm.DB) (gormDB *gorm.DB) {
	gormDB = db
	if value := ctx.Value(constants.GormTransactionKey); value != nil {
		if gormTx, ok := value.(*gorm.DB); ok {
			gormDB = gormTx
		}
	}
	return
}
