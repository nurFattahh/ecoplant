package handler

import (
	"ecoplant/repository"
)

type CartHandler struct {
	Repository repository.CartRepository
}

// "Constructor" for postHandler
func NewCartHandler(repo *repository.CartRepository) CartHandler {
	return CartHandler{*repo}
}
