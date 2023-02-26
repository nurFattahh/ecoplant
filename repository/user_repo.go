package repository

import (
	"ecoplant/entity"
	"ecoplant/model"
	"ecoplant/sdk/crypto"

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
		FullName: model.FullName,
		Username: model.Username,
		Password: hashPassword,
	}

	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, err
	}
	return &user, nil
}
