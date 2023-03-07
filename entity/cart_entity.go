package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	ProductIDs uint    `json:"products"`
	Total      float64 `json:"total"`
}
