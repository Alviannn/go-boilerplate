package repositories_mysql

import (
	"github.com/samber/do/v2"
)

type Base struct {
	Helper   Helper   `do:""`
	Accounts Accounts `do:""`
}

func New(i do.Injector) (*Base, error) {
	return do.InvokeStruct[*Base](i)
}

var Package = do.Package(
	do.Lazy(NewHelper),
	do.Lazy(NewAccounts),
	do.Lazy(New),
)
