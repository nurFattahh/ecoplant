package entity

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	Regency     string    `json:"regency"`
	District    string    `json:"district"`
	NumDonate   int       `json:"num_donate"`
	Wallet      float64   `json:"wallet"`
	RemainDay   int       `json:"remain_day"`
	Plan        string    `json:"plan"`
	News        string    `json:"news"`
	Community   Community `gorm:"foreignKey:CommunityID" json:"community"`
	CommunityID uint      `json:"community_id"`
}
