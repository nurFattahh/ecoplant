package handler

import (
	"ecoplant/model"
	"ecoplant/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	Repository repository.CartRepository
}

// "Constructor" for postHandler
func NewCartHandler(repo *repository.CartRepository) CartHandler {
	return CartHandler{*repo}
}

func (h *CartHandler) AddItem(c *gin.Context) {
	var req struct {
		ItemID uint `json:"item_id"`
		Qty    int  `json:"qty"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.Repository.GetProductByID(req.ItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartItem := &model.CartItem{Product: *item, Quantity: req.Qty}

	if err := h.Repository.AddItem(cartItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartItem)
}

// func (h *CartHandler) CreateCart(c *gin.Context) {
// 	// bind incoming http request
// 	request := model.CreatePostRequest{}
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		response.FailOrError(c, http.StatusUnprocessableEntity, "Create post failed", err)
// 		return
// 	}

// 	result, exist := c.Get("user")
// 	if !exist {
// 		response.FailOrError(c, http.StatusUnprocessableEntity, "no user key found", errors.New("erorr"))
// 		return
// 	}
// 	claims, ok := result.(jwt.MapClaims)
// 	if !ok {
// 		response.FailOrError(c, http.StatusUnprocessableEntity, "error parsing ", errors.New("erorr"))
// 		return
// 	}

// 	userIDc := claims["id"]
// 	userIDf, ok := userIDc.(float64)
// 	if !ok {
// 		response.FailOrError(c, http.StatusUnprocessableEntity, "error get id", errors.New("erorr"))
// 		return
// 	}

// 	// create post
// 	post := entity.Post{
// 		Title:   request.Title,
// 		Content: request.Content,
// 		UserID:  uint(userIDf),
// 	}
// 	err := h.Repository.CreatePost(&post)
// 	if err != nil {
// 		response.FailOrError(c, http.StatusInternalServerError, "Create post failed", err)
// 		return
// 	}
// }
