package entity

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	Name        string  `json:"name"`
	Domisili    string  `json:"domisili"`
	Description string  `json:"description"`
	Wallet      float64 `json:"wallet"`
	User        []User  `gorm:"many2many:user_community;" json:"-"`
}
