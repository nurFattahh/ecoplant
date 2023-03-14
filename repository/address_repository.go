package repository

import (
	"ecoplant/entity"

	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return AddressRepository{db}
}

func (r *AddressRepository) CreateAdress(address *entity.ShippingAddress) error {
	return r.db.Create(address).Error
}

func (r *AddressRepository) ShippingAddress(ID uint, ShippingAddress *entity.ShippingAddress) error {
	return r.db.Where("id = ?", ID).Updates(&ShippingAddress).Error
}
