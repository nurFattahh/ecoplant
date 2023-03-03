package model

type RegisterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	UsernameOrEmail string `json:"username/email"`
	Password        string `json:"password"`
}
