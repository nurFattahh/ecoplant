package repository

import (
	"ecoplant/entity"
	"ecoplant/model"
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

func (r *UserRepository) CreateUser(model model.RegisterUser) (*entity.User, error) {
	hashPassword, err := crypto.HashValue(model.Password)

	if err != nil {
		return nil, err
	}

	var user entity.User = entity.User{
		Name:     model.Name,
		Email:    model.Email,
		Username: model.Username,
		Password: hashPassword,
	}

	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, errors.New("USERNAME ALREADY IN USE")
	}
	return &user, nil
}

func (r *UserRepository) FindByUsernameOrEmail(UsernameOrEmail string) (entity.User, error) {
	user := entity.User{}
	err := r.db.Where("username = ? or email = ?", UsernameOrEmail, UsernameOrEmail).First(&user).Error
	return user, err
}

func (r *UserRepository) GetUserById(id uint) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("id = ?", id).Preload("Transaction").Preload("Address").Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
