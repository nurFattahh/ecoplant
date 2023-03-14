package entity

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	Name        string     `json:"name"`
	Picture     string     `json:"picture"`
	Email       string     `gorm:"unique" json:"email"`
	Description string     `json:"description"`
	Phone       string     `json:"phone"`
	NumMember   int        `json:"num_member"`
	Activities  []Donation `json:"activites"`
	Document    string     `json:"document"`
}
