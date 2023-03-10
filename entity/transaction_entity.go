package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Product           Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity          int     `json:"quantity"`
	Total             float64 `json:"total"`
	Status            string  `gorm:"default:'sudah dibayar'" json:"status"`
	Address           string  `json:"address"`
	Method            string  `json:"method" binding:"required"`
	UserID            uint    `json:"user_id"`
	ProductID         uint    `json:"product_id"`
	ShippingAddressID uint    `json:"shipping_address_id"`
}
