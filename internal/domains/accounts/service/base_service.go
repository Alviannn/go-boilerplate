package accounts_service

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"
)

type serviceImpl struct {
	MySQLRepository accounts_interfaces.MySQLRepository
}

func NewService(mysqlRepository accounts_interfaces.MySQLRepository) accounts_interfaces.Service {
	return &serviceImpl{
		MySQLRepository: mysqlRepository,
	}
}
