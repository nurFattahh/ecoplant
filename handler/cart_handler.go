package handler

import (
	"ecoplant/entity"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CartHandler struct {
	Repository repository.CartRepository
}

func NewCartHandler(repo *repository.CartRepository) CartHandler {
	return CartHandler{*repo}
}

func (h *CartHandler) AddProductToCart(c *gin.Context) {
	request := entity.AddProduct{}
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

	user, err := h.Repository.GetUserCartId(uint(userIDf))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "failed get user", err)
	}

	product, err := h.Repository.GetProductByID(request.ProductID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed get product", err)
		return
	}

	addProduct := entity.CartItem{
		Product:   *product,
		Total:     float64(product.Price),
		ProductID: request.ProductID,
		CartID:    user.CartID,
	}
	total := 0

	err = h.Repository.AddProductToCart(addProduct.CartID, &addProduct)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed add product", err)
		return
	}

	items, err := h.Repository.GetCheckListProduct(user.CartID)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed get cl", err)
		return
	}

	for _, item := range items {
		total += int(item.Total)
	}

	err = h.Repository.UpdateTotal(user.CartID, float64(total))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed update total", err)
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

func (h *CartHandler) DeleteItemInCartByID(c *gin.Context) {
	query := c.Query("product_id")

	parseQuery, _ := strconv.ParseInt(query, 10, 64)

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

	cart, err := h.Repository.GetUserCartId(uint(userIDf))
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Failed get user cart id", err)
		return
	}

	product, err := h.Repository.GetProductByID(uint(parseQuery))
	price := product.Price

	IDCart := cart.CartID

	err = h.Repository.DeleteItemInCartByID(IDCart, price, uint(parseQuery))
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Failed delete item", err)
		return
	}

	response.Success(c, http.StatusOK, "Success delete Item", err)

}
