package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID  uint      `json:"user_id"`
	Product []Product `gorm:"many2many:cart_product" json:"products"`
}
