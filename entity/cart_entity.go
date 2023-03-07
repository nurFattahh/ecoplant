package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	TransactionID uint      `json:"user_id"`
	Product       []Product `json:"product_id"`
	Total         float64   `json:"total"`
}
