package model

type AddProduct struct {
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Address   string `json:"address"`
	Method    string `json:"method"`
}
