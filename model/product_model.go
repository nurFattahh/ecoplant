package model

import "ecoplant/entity"

type CreateProduct struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}

type PostParam struct {
	PostID int64 `uri:"post_id" gorm:"column:id"`
	entity.PaginationParam
}
