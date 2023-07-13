package dtos

type RegisterUserReq struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserReq struct {
	ID int64 `param:"id"`
}

type GetAllUsersReq struct {
	Username string `query:"username" json:"username"`
	Email    string `query:"email" json:"email"`
	FullName string `query:"fullName" json:"fullName"`

	Limit  int `query:"limit" json:"limit"`
	Offset int `query:"offset" json:"offset"`
}
