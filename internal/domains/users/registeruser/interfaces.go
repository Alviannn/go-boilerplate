package registeruser

import "go-boilerplate/internal/dtos"

type Service interface {
	RegisterUser(params dtos.RegisterUserReq) (err error)
}

type Repository interface {
	RegisterUser(params dtos.RegisterUserReq) error
	IsUserExistByEmail(email string) bool
}
