package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"type:VARCHAR(50); NOT NULL" json:"fullname" `
	Username string `gorm:"type:VARCHAR(50); NOT NULL" json:"username" `
	Email    string `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password string `gorm:"type:TEXT; NOT NULL" json:"-" `
	Address  string `gorm:"type:VARCHAR(50)" json:"address"`
	Phone    string `gorm:"type:VARCHAR(50); NOT NULL" json:"phone"`
}
