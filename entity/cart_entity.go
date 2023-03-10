package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Product     Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity    int     `json:"quantity"`
	Total       float64 `json:"total"`
	UserID      uint    `json:"user_id"`
	ProductID   uint    `json:"product_id"`
	IsCheckList bool    `gorm:"default:true" json:"checklist"`
}
