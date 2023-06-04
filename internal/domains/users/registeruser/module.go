package registeruser

import "github.com/goava/di"

var Module = di.Options(
	di.Provide(NewRepository),
	di.Provide(NewService),
	di.Provide(NewHandler),
)
