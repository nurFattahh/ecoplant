package model

type CreateProduct struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	Merchant    string `json:"merchant"`
	Picture     string `gorm:"type:VARCHAR(255)" json:"picture"`
}
