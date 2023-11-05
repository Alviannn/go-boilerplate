package accounts_service

import (
	"context"

	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

func (s *serviceImpl) GetByID(ctx context.Context, params dtos.AccountGetRequest) (account models_mysql.Account, err error) {
	account, err = s.RepositoryMySQL.GetByID(ctx, params.ID)
	return
}
