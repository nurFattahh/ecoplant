package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name              string          `gorm:"type:VARCHAR(50); NOT NULL" json:"name" `
	Username          string          `gorm:"type:VARCHAR(50); uniqueIndex; NOT NULL" json:"username" `
	Email             string          `gorm:"type:VARCHAR(50); NOT NULL" json:"email"`
	Password          string          `gorm:"type:TEXT; NOT NULL" json:"-" `
	Address           ShippingAddress `gorm:"foreignKey:ShippingAddressID" json:"address"`
	Cart              []Cart          `json:"-"`
	Transaction       []Transaction   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transaction"`
	ShippingAddressID uint            `json:"-"`
}

type RegisterUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	UsernameOrEmail string `json:"username/email"`
	Password        string `json:"password"`
}
