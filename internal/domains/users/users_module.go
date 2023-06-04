package users

import (
	"go-boilerplate/internal/domains/users/getuser"
	"go-boilerplate/internal/domains/users/registeruser"

	"github.com/goava/di"
)

var Module = di.Options(
	getuser.Module,
	registeruser.Module,
)
