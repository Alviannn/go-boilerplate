package dtos

type AccountRegisterReq struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AccountGetReq struct {
	ID int64 `param:"id" validate:"required"`
}

type AccountGetAllReq struct {
	Username string `query:"username" json:"username"`
	Email    string `query:"email" json:"email"`
	FullName string `query:"fullName" json:"fullName"`

	Limit  int `query:"limit" json:"limit"`
	Offset int `query:"offset" json:"offset"`
}
