package accounts_repository_mysql

import (
	"go-boilerplate/internal/constants"
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"
	"go-boilerplate/internal/dtos"

	"gorm.io/gorm"
)

func (r *repositoryImpl) Transaction(deps dtos.TxDependencies) accounts_interfaces.RepositoryMySQL {
	return &repositoryImpl{
		Tx: deps[constants.DependencyKeyMySQLTx].(*gorm.DB),
	}
}
