package model

type RegisterUser struct {
	FullName string
	Username string
	Password string
}

type LoginUser struct {
	Username string
	Password string
}
