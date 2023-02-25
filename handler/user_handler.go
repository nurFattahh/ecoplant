package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type postHandler struct {
	DB *gorm.DB
}

// "Constructor" for postHandler
func NewPostHandler(db *gorm.DB) postHandler {
	return postHandler{db}
}

func (h *postHandler) CreatePost(c *gin.Context) {
	// ini diisi nanti
}
