package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	Merchant    string `json:"merchant"`
	Picture     string `gorm:"type:VARCHAR(255)" json:"picture"`
}
