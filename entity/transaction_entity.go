package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID uint `json:"user_id"`
	CartID uint `json:"cart_id"`
}
