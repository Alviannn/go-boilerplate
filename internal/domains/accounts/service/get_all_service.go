package accounts_service

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

func (s *serviceImpl) GetAll(ctx context.Context, params dtos.AccountGetAllRequest) (accounts []models_mysql.Account, err error) {
	accounts, err = s.RepositoryMySQL.GetAll(ctx, params)
	return
}
