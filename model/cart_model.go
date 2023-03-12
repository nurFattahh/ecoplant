package model

type AddProduct struct {
	ProductID   uint `json:"product_id"`
	Quantity    int  `json:"quantity"`
	IsCheckList bool `json:"is_checklist"`
}
