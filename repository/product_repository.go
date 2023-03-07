package repository

import (
	"ecoplant/entity"
	"ecoplant/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func (r *ProductRepository) GetAllProduct(model *model.PaginParam) ([]entity.Product, int, error) {
	var products []entity.Product
	err := r.db.
		Model(entity.Product{}).
		Limit(model.Limit).
		Offset(model.Offset).
		Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	var totalElements int64
	err = r.db.
		Model(entity.Product{}).
		Limit(model.Limit).
		Offset(model.Offset).
		Count(&totalElements).Error
	if err != nil {
		return nil, 0, err
	}
	return products, int(totalElements), err
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

func (h *ProductRepository) BindBody(c *gin.Context, body interface{}) interface{} {
	return c.ShouldBindWith(body, binding.JSON)
}

func (h *ProductRepository) BindParam(c *gin.Context, param interface{}) error {
	if err := c.ShouldBindUri(param); err != nil {
		return err
	}
	return c.ShouldBindWith(param, binding.Query)
}
