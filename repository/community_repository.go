package repository

import (
	"ecoplant/entity"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type CommunityRepository struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) CommunityRepository {
	return CommunityRepository{db}
}

func (r *CommunityRepository) CreateCommunity(model *entity.Community) error {
	return r.db.Create(model).Error
}

func (r *CommunityRepository) GetCommunityByName(query string) (*[]entity.Community, error) {
	var community []entity.Community
	result := r.db.Where("name LIKE ?", "%"+query+"%").Preload("Activities").Find(&community)

	if result.Error != nil {
		return nil, result.Error
	}
	return &community, nil
}

func (r *CommunityRepository) GetAllCommunity(model *entity.PaginParam) ([]entity.Community, int, error) {
	var community []entity.Community
	err := r.db.
		Model(entity.Community{}).
		Limit(model.Limit).
		Offset(model.Offset).
		Preload("Activities").
		Find(&community).Error
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
	return community, int(totalElements), err
}

func (r *CommunityRepository) GetCommunityByID(ID uint) (*entity.Community, error) {
	var community entity.Community

	result := r.db.Preload("Activities").First(&community, ID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &community, nil
}

func (h *CommunityRepository) BindBody(c *gin.Context, body interface{}) interface{} {
	return c.ShouldBindWith(body, binding.JSON)
}

func (h *CommunityRepository) BindParam(c *gin.Context, param interface{}) error {
	if err := c.ShouldBindUri(param); err != nil {
		return err
	}
	return c.ShouldBindWith(param, binding.Query)
}
