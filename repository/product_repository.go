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

func (r *ProductRepository) CreateProduct(post *entity.Product) error {
	return r.db.Create(post).Error
}

func (r *ProductRepository) GetAllProduct() ([]entity.Product, error) {
	var posts []entity.Product

	err := r.db.Find(&posts).Error

	return posts, err
}

func (r *ProductRepository) GetProductByID(ID uint) (*entity.Product, error) {
	var product entity.Product

	result := r.db.First(&product, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *ProductRepository) GetProductByName(name string) (*[]entity.Product, error) {
	var products []entity.Product
	result := r.db.Where("name LIKE ?", "%"+name+"%").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil

}

func (r *ProductRepository) UpdateProduct(ID uint, updateProduct *model.UpdateProductRequest) error {
	var product entity.Product

	err := r.db.Model(&product).Where("id = ?", ID).Updates(updateProduct).Error

	return err
}
