package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Product  []Cart `gorm:"many2many:transaction_cart" json:"product"`
	Quantity int    `json:"quantity"`
}
