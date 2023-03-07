package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string        `gorm:"type:VARCHAR(50); NOT NULL" json:"name" `
	Username    string        `gorm:"type:VARCHAR(50); uniqueIndex; NOT NULL" json:"username" `
	Email       string        `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password    string        `gorm:"type:TEXT; NOT NULL" json:"-" `
	Address     Address       `gorm:"type:VARCHAR(50)" json:"address"`
	Phone       string        `gorm:"type:VARCHAR(50)" json:"phone"`
	Transaction []Transaction `json:"transaction"`
}
