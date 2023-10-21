package accounts_postgres_repository

import (
	"context"
	"go-boilerplate/internal/dtos"
	postgres_models "go-boilerplate/internal/models/postgres"
)

func (r *postgresRepositoryImpl) GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []postgres_models.Account, err error) {
	query := r.DB.WithContext(ctx)

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

	err = query.Find(&accounts).Error
	return
}
