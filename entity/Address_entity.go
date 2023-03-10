package entity

type ShippingAddress struct {
	ShippingAddressID uint   `gorm:"autoIncrement;unique" json:"id"`
	Recipient         string `json:"recipient"`
	Phone             string `json:"phone"`
	Province          string `json:"province"`
	RegencyDistrict   string `json:"regency_district"`
	Home              string `json:"home"`
	PostalCode        uint   `json:"postal_code"`
}
