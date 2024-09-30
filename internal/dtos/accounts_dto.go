package dtos

type AccountRegisterReq struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountGetReq struct {
	ID int64 `param:"id"`
}

type AccountGetAllReq struct {
	Username string `query:"username" json:"username"`
	Email    string `query:"email" json:"email"`
	FullName string `query:"fullName" json:"fullName"`

	Limit  int `query:"limit" json:"limit"`
	Offset int `query:"offset" json:"offset"`
}
