package services

import "github.com/samber/do/v2"

type Base struct {
	Accounts Accounts `do:""`
	Helper   Helper   `do:""`
}

func New(i do.Injector) (*Base, error) {
	return do.InvokeStruct[*Base](i)
}

var Package = do.Package(
	do.Lazy(New),
	do.Lazy(NewHelper),
	do.Lazy(NewAccounts),
)
