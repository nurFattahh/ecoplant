package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Price       int
	Rating      int
	Description string
}
