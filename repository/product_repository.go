package repository

import (
	"ecoplant/entity"
	"ecoplant/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (r *ProductRepository) CreateProduct(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetAllProduct() ([]entity.Product, error) {
	var products []entity.Product

	err := r.db.Find(&products).Error

	return products, err
}

func (r *ProductRepository) GetProductByID(ID uint) (*entity.Product, error) {
	var product entity.Product

	result := r.db.First(&product, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *ProductRepository) GetProductByName(query string) (*[]entity.Product, error) {
	var products []entity.Product
	result := r.db.Where("name LIKE ?", "%"+query+"%").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil

}

func (r *ProductRepository) UpdateProduct(ID uint, updatePost *model.UpdateProduct) error {
	var product entity.Product

	err := r.db.Model(&product).Where("id = ?", ID).Updates(updatePost).Error

	return err
}

func (r *ProductRepository) DeleteProduct(ID uint) error {
	var product entity.Product

	err := r.db.Delete(&product, ID).Error

	return err
}
