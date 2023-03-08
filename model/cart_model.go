package model

type AddProduct struct {
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

// type CartItem struct {
// 	Product  entity.Product `json:"product"`
// 	Quantity int            `json:"quantity"`
// }
