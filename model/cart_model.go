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

// type CartItem struct {
// 	Product  entity.Product `json:"product"`
// 	Quantity int            `json:"quantity"`
// }
