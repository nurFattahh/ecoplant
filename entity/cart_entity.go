package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Items  []CartItem `json:"items"`
	UserID uint       `json:"user_id"`
	Total  float64    `json:"total"`
}

type CartItem struct {
	Product     Product `gorm:"foreignKey:ProductID" json:"product"`
	IsCheckList bool    `gorm:"default:false" json:"checklist"`
	Quantity    int     `json:"quantity"`
	Total       float64 `json:"total"`
	ProductID   uint    `json:"product_id"`
	CartID      uint    `json:"cart_id"`
}

type AddProduct struct {
	ProductID   uint `json:"product_id"`
	Quantity    int  `json:"quantity"`
	IsCheckList bool `json:"is_checklist"`
}
