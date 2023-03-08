package model

type Checkout struct {
	Quantity  int  `json:"quantity"`
	ProductID uint `json:"product_id"`
}
