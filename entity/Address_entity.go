package entity

import "gorm.io/gorm"

type ShippingAddress struct {
	gorm.Model
	Recipient       string `json:"recipient"`
	Phone           string `json:"phone"`
	Province        string `json:"province"`
	RegencyDistrict string `json:"regency_district"`
	Home            string `json:"home"`
	PostalCode      int    `json:"postal_code"`
	UserID          uint   `json:"-"`
}

type GetAddress struct {
	Recipient       string `json:"recipient"`
	Phone           string `json:"phone"`
	Province        string `json:"province"`
	RegencyDistrict string `json:"regency_district"`
	Home            string `json:"home"`
	PostalCode      int    `json:"postal_code"`
}
