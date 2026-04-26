package main

import (
	repositories_mysql "go-boilerplate/internal/repositories/mysql"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/customvalidator"
	"go-boilerplate/pkg/databases"

	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func NewDI() *do.RootScope {
	injector := do.New()

	do.Provide(injector, func(i do.Injector) (*customvalidator.Validator, error) {
		return customvalidator.New(), nil
	})
	do.Provide(injector, func(i do.Injector) (*gorm.DB, error) {
		return databases.NewMySQLDB()
	})

	repositories_mysql.Package(injector)
	services.Package(injector)

	return injector
}
