package repository

import (
	"ecoplant/entity"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (r *CartRepository) GetProductByID(ID uint) (*entity.Product, error) {
	var product entity.Product

	result := r.db.First(&product, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *CartRepository) AddProductToCart(input *entity.Cart) error {
	return r.db.Create(&input).Error
}

func (r *CartRepository) GetAllProductInCart(ID uint) ([]entity.Cart, error) {
	var cart []entity.Cart
	err := r.db.Model(entity.Cart{}).Preload("Product").Find(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}
