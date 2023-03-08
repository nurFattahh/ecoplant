package repository

// import (
// 	"ecoplant/entity"

// 	"gorm.io/gorm"
// )

// type CartRepository struct {
// 	db *gorm.DB
// }

// func NewCartRepository(db *gorm.DB) CartRepository {
// 	return CartRepository{db}
// }

// func (r *CartRepository) GetProductByID(ID uint) (*entity.Product, error) {
// 	var product entity.Product

// 	result := r.db.First(&product, ID)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &product, nil
// }

// func (r *CartRepository) CreateTransaction(transaction *entity.Transaction) error {
// 	return r.db.Create(transaction).Error
// }

// func (r *CartRepository) GetCart(userID int, produkID int) error {
// 	product, err := r.GetProductByID(uint(produkID))
// 	if err != nil {
// 		return err
// 	}

// 	var carts []entity.Cart
// 	// carts = append(carts, entity.Cart{UserID: uint(userID), Product: *product})
// 	return nil
// }
