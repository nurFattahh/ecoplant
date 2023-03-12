package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Product           Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity          int     `json:"quantity"`
	TotalProduct      float64 `json:"total_product"`
	Status            string  `gorm:"default:'sudah dibayar'" json:"status"`
	Address           string  `json:"address"`
	PaymentMethod     string  `json:"payment_method"`
	PaymentPrice      float64 `json:"payment_price"`
	ShippingMethod    string  `json:"shipping_method"`
	Estimate          string  `json:"estimate"`
	UserID            uint    `json:"user_id"`
	ProductID         uint    `json:"product_id"`
	ShippingAddressID uint    `json:"-"`
	Total             float64 `json:"total"`
}
