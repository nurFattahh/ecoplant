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

type AddressHandler struct {
	Repository repository.AddressRepository
}

func NewAddressHandler(repo *repository.AddressRepository) AddressHandler {
	return AddressHandler{*repo}
}

func (h *AddressHandler) ShippingAddress(c *gin.Context) {
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

	userAddress := claims["shipping"]
	userAddressc, ok := userAddress.(float64)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error get shipping", errors.New("erorr"))
		return
	}

	address := entity.ShippingAddress{
		Recipient:       request.Recipient,
		Phone:           request.Phone,
		Province:        request.Province,
		RegencyDistrict: request.RegencyDistrict,
		Home:            request.Home,
		PostalCode:      request.PostalCode,
		UserID:          uint(userIDf),
	}

	err := h.Repository.ShippingAddress(uint(userAddressc), &address)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed Post Address", err)
		return
	}

	response.Success(c, http.StatusOK, "Success Post Address", nil)

}
