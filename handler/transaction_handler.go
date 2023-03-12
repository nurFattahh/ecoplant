package handler

import (
	"ecoplant/entity"
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

	request := entity.Checkout{}
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

	meth := request.PaymentMethod
	var method string
	switch meth {
	case 1:
		method = "Bank BCA"
	case 2:
		method = "Bank BRI"
	case 3:
		method = "Bank Mandiri"
	}

	shipping := request.ShippingMethod
	var shippingMethod string
	var paymentPrice float64
	var estimate string

	switch shipping {
	case 1:
		shippingMethod = "JNE Regular"
		paymentPrice = 12000
		estimate = "3 - 5 hari"
	case 2:
		shippingMethod = "J&T Express"
		paymentPrice = 20000
		estimate = "3 - 5 hari"
	case 3:
		shippingMethod = "Sicepat Ekonomi"
		paymentPrice = 18000
		estimate = "3 - 5 hari"
	}
	TotalProduct := request.Quantity * product.Price
	total := request.Quantity*product.Price + int(paymentPrice)

	transaction := entity.Transaction{
		Product:        *product,
		Quantity:       request.Quantity,
		TotalProduct:   float64(TotalProduct),
		Total:          float64(total),
		Address:        address.RegencyDistrict,
		ShippingMethod: shippingMethod,
		PaymentMethod:  method,
		PaymentPrice:   paymentPrice,
		Estimate:       estimate,
		Status:         request.Status,
		UserID:         uint(userIDf),
		ProductID:      request.ProductID,
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
	request := entity.GetAddress{}
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
