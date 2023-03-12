package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Product     Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity    int     `json:"quantity"`
	Total       float64 `json:"total"`
	IsCheckList bool    `gorm:"default:false" json:"checklist"`
	UserID      uint    `json:"user_id"`
	ProductID   uint    `json:"product_id"`
}
