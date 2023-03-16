package entity

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	Name        string    `json:"name"`
	Regency     string    `json:"regency"`
	District    string    `json:"district"`
	NumDonate   int       `json:"num_donate"`
	Wallet      float64   `json:"wallet"`
	Target      float64   `json:"target"`
	RemainDay   int       `json:"remain_day"`
	Plan        string    `json:"plan"`
	News        string    `json:"news"`
	Picture     string    `json:"picture"`
	Community   Community `gorm:"foreignKey:CommunityID" json:"-"`
	CommunityID uint      `json:"-"`
}

type UserDonation struct {
	gorm.Model
	UserID        uint     `json:"user_id"`
	DonationID    uint     `json:"donation_id"`
	Nominal       float64  `json:"nominal"`
	PaymentMethod string   `json:"payment_method"`
	Status        string   `gorm:"default:'Donasi Selesai'" json:"status"`
	Donation      Donation `gorm:"foreignKey:DonationID" json:"-"`
}

type CreateDonation struct {
	Picture   string  `json:"picture"`
	Name      string  `json:"name"`
	Regency   string  `json:"regency"`
	District  string  `json:"district"`
	Target    float64 `json:"target"`
	RemainDay int     `json:"remain_day"`
	Plan      string  `json:"plan"`
	News      string  `json:"news"`
}

type UserDonationRequest struct {
	Nominal       float64 `json:"nominal"`
	PaymentMethod int     `json:"payment_method"`
}

type UpdatePlanAndNewsDonation struct {
	Plan string `json:"plan"`
	News string `json:"news"`
}
