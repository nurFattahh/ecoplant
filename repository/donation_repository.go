package repository

import (
	"ecoplant/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type DonationRepository struct {
	db *gorm.DB
}

func NewDonationRepository(db *gorm.DB) DonationRepository {
	return DonationRepository{db}
}

func (r *DonationRepository) CreateDonation(model *entity.Donation) error {
	return r.db.Create(model).Error
}

func (r *DonationRepository) GetDonationByRegency(query string) (*[]entity.Donation, error) {
	var donation []entity.Donation
	result := r.db.Where("regency LIKE ?", "%"+query+"%").Find(&donation)

	if result.Error != nil {
		return nil, result.Error
	}
	return &donation, nil
}

func (r *DonationRepository) GetAllDonation(model *entity.PaginParam) ([]entity.Donation, int, error) {
	var donation []entity.Donation
	err := r.db.
		Model(entity.Donation{}).
		Limit(model.Limit).
		Offset(model.Offset).
		Find(&donation).Error
	if err != nil {
		return nil, 0, err
	}
	var totalElements int64
	err = r.db.
		Model(entity.Donation{}).
		Limit(model.Limit).
		Offset(model.Offset).
		Count(&totalElements).Error
	if err != nil {
		return nil, 0, err
	}
	return donation, int(totalElements), err
}

func (r *DonationRepository) GetDonationByID(ID uint) (*entity.Donation, error) {
	var donation entity.Donation

	result := r.db.First(&donation, ID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &donation, nil
}

func (r *DonationRepository) GetCommunityByID(ID uint) (*entity.Community, error) {
	var community entity.Community

	result := r.db.First(&community, ID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &community, nil
}

func (r *DonationRepository) CreateUserDonation(donationID, nominal float64, model *entity.UserDonation) error {
	err := r.db.Model(entity.Donation{}).Where("id =?", donationID).Update("wallet", gorm.Expr("wallet + ?", nominal)).Error
	err = r.db.Model(entity.Donation{}).Where("id =?", donationID).Update("num_donate", gorm.Expr("num_donate + ?", 1)).Error
	if err != nil {
		return err
	}
	return r.db.Create(model).Error
}

func (r *DonationRepository) GetAllUserDonation(ID uint) ([]entity.UserDonation, error) {
	var donation []entity.UserDonation
	err := r.db.Model(entity.UserDonation{}).Where("user_id =?", ID).Find(&donation).Error
	if err != nil {
		return nil, err
	}
	return donation, nil
}

func (r *DonationRepository) UpdatePlanAndNewsDonation(id uint, model entity.Donation) error {
	err := r.db.Model(entity.Donation{}).Where("id =?", id).Updates(model).Error
	return err
}

func (h *DonationRepository) BindBody(c *gin.Context, body interface{}) interface{} {
	return c.ShouldBindWith(body, binding.JSON)
}

func (h *DonationRepository) BindParam(c *gin.Context, param interface{}) error {
	if err := c.ShouldBindUri(param); err != nil {
		return err
	}
	return c.ShouldBindWith(param, binding.Query)
}
