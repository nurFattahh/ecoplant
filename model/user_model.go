package model

import "gorm.io/gorm"

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

type ResponseRegister struct {
	gorm.Model
	Name     string `gorm:"type:VARCHAR(50); NOT NULL" json:"name" `
	Username string `gorm:"type:VARCHAR(50); uniqueIndex; NOT NULL" json:"username" `
	Email    string `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password string `gorm:"type:TEXT; NOT NULL" json:"-" `
	Phone    string `gorm:"type:VARCHAR(50)" json:"phone"`
}
