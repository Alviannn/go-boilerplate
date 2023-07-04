package dependencies

import (
	"go-boilerplate/internal/domains"

	"github.com/goava/di"
	"gorm.io/gorm"
)

// New creates a new DI (dependency injection) container
// with the modules we provided in `modules.go`.
func New() (container *di.Container, err error) {
	return di.New(
		appModules,
		domains.Modules,
	)
}

// NewForTransaction creates a new DI container just like `New`
// but specifically used for database transaction.
//
// This will be only providing the dependencies from 'domains'
// directory, meaning only the features will be re-instantiated
// and not the full container from `New` to only get the necessary
// dependencies. Don't add any extra dependencies to this function,
// and instead use `container.ProvideValue`.
//
// Example using for database transaction:
//
//	tx := db.Begin()
//	txContainer, err := dependencies.NewForTransaction(tx)
//	if err != nil {
//		return err
//	}
//
//	// For example, the accounts repository requires 'customvalidator'
//	// or else it will throw an error.
//	if err = txContainer.ProvideValue(validator); err != nil {
//		return err
//	}
//
//	var accountsRepo accounts_interfaces.Repository
//	if err = txContainer.Resolve(&accountsRepo); err != nil {
//		return err
//	}
//
//	// Use the resolved instance (in this case it's repo) for transactions.
//	if err = accountsRepo.GetAccount(1); err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	if err = tx.Commit().Error; err != nil {
//		return err
//	}
func NewForTransaction(db *gorm.DB) (txContainer *di.Container, err error) {
	return di.New(
		di.ProvideValue(db),
		domains.Modules,
	)
}
