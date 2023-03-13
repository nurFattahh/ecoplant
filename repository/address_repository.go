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

func (r *AddressRepository) GetAddress(id uint) (*entity.ShippingAddress, error) {
	var address entity.ShippingAddress
	result := r.db.Where("user_id = ?", id).Take(&address)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}
