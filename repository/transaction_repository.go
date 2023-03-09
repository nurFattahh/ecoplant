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
