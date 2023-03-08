package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID   uint    `json:"user_id"`
	Product  Product `gorm:"many2many:transaction_cart" json:"product"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
