package model

import "math"

type CreateProduct struct {
	Name        string `json:"name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Rating      int    `json:"rating" binding:"required"`
	Description string `json:"description" binding:"required"`
	Merchant    string `json:"merchant" binding:"required"`
	Picture     string `gorm:"type:VARCHAR(255)" json:"picture" binding:"required"`
}

type UpdateProduct struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	Merchant    string `json:"merchant"`
	Picture     string `gorm:"type:VARCHAR(255)" json:"picture"`
}

type PaginParam struct {
	Limit           int `json:"limit" form:"limit"`
	Page            int `json:"page" form:"page"`
	Offset          int `json:"offset"`
	TotalElements   int `json:"total_elements"`
	CurrentElements int `json:"current_elements"`
	TotalPages      int `json:"total_pages"`
	CurrentPage     int `json:"current_page"`
}

type UpdateLocation struct {
	Location string `gorm:"location"`
}

func (pp *PaginParam) FormatPagin() {
	if pp.Limit == 0 {
		pp.Limit = 9
	}
	if pp.Page == 0 {
		pp.Page = 1
	}
	pp.Offset = (pp.Page - 1) * pp.Limit
}

func (pp *PaginParam) ProcessPagin(totalElements int) {
	pp.TotalElements = totalElements
	pp.TotalPages = int(math.Ceil(float64(pp.TotalElements) / float64(pp.Limit)))
	pp.CurrentPage = pp.Page
	if pp.TotalPages > pp.TotalElements/pp.Limit {
		if pp.Page < pp.TotalPages {
			pp.CurrentElements = pp.Limit
		} else {
			pp.CurrentElements = pp.TotalElements - (pp.TotalPages-1)*pp.Limit
		}
	}
}
