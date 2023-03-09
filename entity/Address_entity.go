package entity

type Address struct {
	Country    string `json:"country"`
	Province   string `json:"province"`
	Regency    string `json:"regency"`
	District   string `json:"district"`
	Home       string `json:"home"`
	PostalCode string `json:"postal"`
	UserID     uint   `json:"-"`
}
