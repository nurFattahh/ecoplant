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

type TransactionHandler struct {
	Repository repository.TransactionRepository
}

func NewTransactionHandler(repo *repository.TransactionRepository) TransactionHandler {
	return TransactionHandler{*repo}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {

	request := model.Checkout{}
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

	product, err := h.Repository.GetProductByID(uint(request.ProductID))
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed get product", err)
	}

	address, err := h.Repository.GetAddress(uint(userIDf))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed getting Address", err)
		return
	}

	total := request.Quantity * product.Price

	transaction := entity.Transaction{
		Product:   *product,
		Quantity:  request.Quantity,
		Total:     float64(total),
		Address:   *&address.RegencyDistrict,
		Method:    request.Method,
		Status:    request.Status,
		UserID:    uint(userIDf),
		ProductID: request.ProductID,
	}

	err = h.Repository.CreateTransaction(&transaction)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Transaction Failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "Transaction Success", transaction)
}

func (h *TransactionHandler) GetAllTransactionByBearer(c *gin.Context) {
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
	transaction, err := h.Repository.GetAllTransactionByBearer(uint(userIDf))

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Transaction not found", err)
		return
	}

	response.Success(c, http.StatusOK, "Transaction Found", transaction)
}

func (h *TransactionHandler) ShippingAddress(c *gin.Context) {
	request := model.GetAddress{}
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
	address := entity.ShippingAddress{
		ShippingAddressID: uint(userIDf),
		Recipient:         request.Recipient,
		Phone:             request.Phone,
		Province:          request.Province,
		RegencyDistrict:   request.RegencyDistrict,
		Home:              request.Home,
		PostalCode:        request.PostalCode,
	}

	err := h.Repository.ShippingAddress(uint(userIDf), &address)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed Post Address", err)
		return
	}

	response.Success(c, http.StatusOK, "Success Post Address", nil)

}
