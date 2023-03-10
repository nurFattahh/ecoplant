package handler

import (
	"ecoplant/repository"
)

type CartHandler struct {
	Repository repository.CartRepository
}

func NewCartHandler(repo *repository.CartRepository) CartHandler {
	return CartHandler{*repo}
}

// func (h *CartHandler) AddProductToCart(c *gin.Context) {

// 	request := model.AddProduct{}
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		response.FailOrError(c, http.StatusUnprocessableEntity, "Create transaction failed", err)
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

// 	product, err := h.Repository.GetProductByID(uint(request.ProductID))
// 	if err != nil {
// 		response.FailOrError(c, http.StatusBadRequest, "Failed get product", err)
// 	}

// 	total := request.Quantity * product.Price

// 	AddProduct := entity.Cart{
// 		Product: *request.Product,
// 		UserID:  uint(userIDf),
// 	}

// 	err = h.Repository.CreateTransaction(&transaction)
// 	if err != nil {
// 		response.FailOrError(c, http.StatusInternalServerError, "Transaction Failed", err)
// 		return
// 	}

// 	response.Success(c, http.StatusCreated, "Transaction Success", transaction)
// }
