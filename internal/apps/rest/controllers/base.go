package controllers_rest

import "github.com/defval/di"

func Module() di.Option {
	return di.Options(
		di.Provide(NewHealth, di.As(new(Controller))),
		di.Provide(NewAccounts, di.As(new(Controller))),
	)
}
