package repository

import (
	"ecoplant/entity"
	"ecoplant/sdk/crypto"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (r *UserRepository) CreateUser(model entity.RegisterUser) (*entity.User, error) {
	hashPassword, err := crypto.HashValue(model.Password)

	if err != nil {
		return nil, err
	}

	var address entity.ShippingAddress = entity.ShippingAddress{}
	r.db.Create(&address)

	var user entity.User = entity.User{
		Name:              model.Name,
		Email:             model.Email,
		Username:          model.Username,
		Password:          hashPassword,
		ShippingAddressID: address.ID,
	}

	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, errors.New("USERNAME ATAU EMAIL SUDAH DIGUNAKAN")
	}

	var cart entity.Cart = entity.Cart{
		UserID: user.ID,
	}
	r.db.Create(&cart)

	r.db.Model(&user).Update("cart_id", cart.ID)

	return &user, nil
}

func (r *UserRepository) FindByUsernameOrEmail(UsernameOrEmail string) (entity.User, error) {
	user := entity.User{}
	err := r.db.Where("username = ? or email = ?", UsernameOrEmail, UsernameOrEmail).First(&user).Error
	return user, err
}

func (r *UserRepository) GetUserById(id uint) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("id = ?", id).Preload("Cart.Items.Product").Preload("Transaction.Product").Preload("Address").Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(id uint, model entity.User) error {
	err := r.db.Where("id = ?", id).Updates(model).Error
	return err
}
