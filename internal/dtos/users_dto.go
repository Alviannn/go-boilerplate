package dtos

import "go-boilerplate/internal/models"

type RegisterUserReq struct {
	models.User

	ID string `json:"-"` // ignore inserting ID
}

type GetUserReq struct {
	UserID string `param:"id"`
}
