package entity

type ShippingAddress struct {
	ShippingAddressID uint   `gorm:"primaryKey" json:"shipping_address_id"`
	Recipient         string `json:"recipient"`
	Phone             string `json:"phone"`
	Province          string `json:"province"`
	RegencyDistrict   string `json:"regency_district"`
	Home              string `json:"home"`
	PostalCode        uint   `json:"postal_code"`
	UserID            uint   `json:"user_id"`
}

type GetAddress struct {
	ShippingAddressID uint   `json:"shipping_address_id"`
	Recipient         string `json:"recipient"`
	Phone             string `json:"phone"`
	Province          string `json:"province"`
	RegencyDistrict   string `json:"regency_district"`
	Home              string `json:"home"`
	PostalCode        uint   `json:"postal_code"`
}
