package services

import "github.com/defval/di"

func Module() di.Option {
	return di.Options(
		di.Provide(NewAccounts),
	)
}
