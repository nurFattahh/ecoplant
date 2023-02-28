package repository

import (
	"ecoplant/entity"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (r *ProductRepository) CreatePost(post *entity.Product) error {
	return r.db.Create(post).Error
}

func (r *ProductRepository) GetAllProduct() ([]entity.Product, error) {
	var posts []entity.Product

	err := r.db.Find(&posts).Error

	return posts, err
}
