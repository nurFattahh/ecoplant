package model

type Checkout struct {
	Quantity       int    `json:"quantity"`
	ProductID      uint   `json:"product_id"`
	Address        string `json:"address"`
	PaymentMethod  int    `json:"payment_method"`
	ShippingMethod int    `json:"shipping_method"`
	Status         string `json:"status"`
}

type GetAddress struct {
	ShippingAddressID uint   `json:"shipping_id"`
	Recipient         string `json:"recipient"`
	Phone             string `json:"phone"`
	Province          string `json:"province"`
	RegencyDistrict   string `json:"regency_district"`
	Home              string `json:"home"`
	PostalCode        uint   `json:"postal_code"`
}
