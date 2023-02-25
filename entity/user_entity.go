package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model // ini sudah mencakup id, dan timestamps
	FullName   string
	Username   string
	Email      string
	Password   string
	Address    string
	Phone      string
}
