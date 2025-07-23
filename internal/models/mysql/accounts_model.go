package models_mysql

type Account struct {
	BaseModel

	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"-"` // don't allow password to ever be exported.
}

func (m *Account) TableName() string {
	return "accounts"
}

type AccountGetParam struct {
	ID       int64
	Username string
	Email    string
}
