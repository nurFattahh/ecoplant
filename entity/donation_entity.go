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
	Community   Community `gorm:"foreignKey:CommunityID" json:"community"`
	CommunityID uint      `json:"community_id"`
}

type CreateDonation struct {
	Name        string  `json:"name"`
	Regency     string  `json:"regency"`
	District    string  `json:"district"`
	Target      float64 `json:"target"`
	RemainDay   int     `json:"remain_day"`
	Plan        string  `json:"plan"`
	News        string  `json:"news"`
	CommunityID uint    `json:"community_id"`
}
