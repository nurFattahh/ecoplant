package entity

type Address struct {
	Country    string `json:"country"`
	Province   string `json:"province"`
	Regency    string `json:"regency"`
	District   string `json:"district"`
	Home       string `json:"home"`
	PostalCode uint   `json:"postal"`
	UserID     uint   `json:"user_id"`
}
