package users

import (
	"go-boilerplate/internal/domains/users/getuser"

	"github.com/goava/di"
)

var Module = di.Options(
	getuser.Module,
)
