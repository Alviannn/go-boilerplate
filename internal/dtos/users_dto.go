package dtos

import "go-boilerplate/internal/models"

type RegisterUserReq struct {
	models.User

	ID string `json:"-"` // ignore inserting ID
}

type GetUserReq struct {
	UserID string `param:"id"`
}

type GetAllUsersReq struct {
	Username string `query:"username"`
	Email    string `query:"email"`
	FullName string `query:"full_name"`

	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}
