package model

type AddProduct struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CreateCartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CreateTransactionRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
