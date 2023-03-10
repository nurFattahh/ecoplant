package repository

import (
	"ecoplant/entity"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{db}
}

func (r *TransactionRepository) GetProductByID(ID uint) (*entity.Product, error) {
	var product entity.Product

	result := r.db.First(&product, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (r *TransactionRepository) CreateTransaction(transaction *entity.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) GetAllTransactionByBearer(user uint) ([]entity.Transaction, error) {
	var transaction []entity.Transaction
	err := r.db.Model(entity.Transaction{}).Preload("Product").Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) ShippingAddress(ID uint, ShippingAddress *entity.ShippingAddress) error {
	// var address entity.ShippingAddress

	err := r.db.Where("shipping_address_id = ?", ID).Updates(&ShippingAddress).Error

	return err
}

func (r *TransactionRepository) GetAddress(id uint) (*entity.ShippingAddress, error) {
	var address entity.ShippingAddress
	result := r.db.Where("shipping_address_id = ?", id).Take(&address)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}
