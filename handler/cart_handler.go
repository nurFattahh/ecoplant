package handler

import (
	"ecoplant/entity"
	"ecoplant/model"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CartHandler struct {
	Repository repository.CartRepository
}

func NewCartHandler(repo *repository.CartRepository) CartHandler {
	return CartHandler{*repo}
}

func (h *CartHandler) CreateProductForCart(c *gin.Context) {
	request := model.AddProduct{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Create transaction failed", err)
		return
	}

	result, exist := c.Get("user")
	if !exist {
		response.FailOrError(c, http.StatusUnprocessableEntity, "no user key found", errors.New("erorr"))
		return
	}
	claims, ok := result.(jwt.MapClaims)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error parsing ", errors.New("erorr"))
		return
	}

	userIDc := claims["id"]
	userIDf, ok := userIDc.(float64)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error get id", errors.New("erorr"))
		return
	}

	product, err := h.Repository.GetProductByID(request.ProductID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed get product", err)
		return
	}

	total := request.Quantity * product.Price

	addProduct := entity.Cart{
		Product:     *product,
		Quantity:    request.Quantity,
		Total:       float64(total),
		UserID:      uint(userIDf),
		IsCheckList: request.IsCheckList,
	}
	err = h.Repository.AddProductToCart(&addProduct)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed add product", err)
		return
	}

	response.Success(c, http.StatusOK, "succes add product", err)

}

func (h *CartHandler) GetAllProductInCart(c *gin.Context) {
	result, exist := c.Get("user")
	if !exist {
		response.FailOrError(c, http.StatusUnprocessableEntity, "no user key found", errors.New("erorr"))
		return
	}
	claims, ok := result.(jwt.MapClaims)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error parsing ", errors.New("erorr"))
		return
	}

	userIDc := claims["id"]
	userIDf, ok := userIDc.(float64)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error get id", errors.New("erorr"))
		return
	}
	carts, err := h.Repository.GetAllProductInCart(uint(userIDf))

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Cart not found", err)
		return
	}

	response.Success(c, http.StatusOK, "Cart Found", carts)
}
