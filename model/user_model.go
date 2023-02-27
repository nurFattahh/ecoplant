package model

type RegisterUser struct {
	FullName string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
