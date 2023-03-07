package repository

import (
	"ecoplant/entity"
	"ecoplant/model"

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

func (r *CartRepository) AddItem(item *model.CartItem) error {
	return r.db.Create(item).Error
}

func (r *CartRepository) AddItemToCart(ID int, item *entity.Cart) error {
	return r.db.Where("id = ?", ID).Create(item).Error
}
