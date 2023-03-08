package handler

// import (
// 	"ecoplant/entity"
// 	"ecoplant/model"
// 	"ecoplant/repository"
// 	"ecoplant/sdk/response"
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// )

// type CartHandler struct {
// 	Repository repository.CartRepository
// }

// // "Constructor" for postHandler
// func NewCartHandler(repo *repository.CartRepository) CartHandler {
// 	return CartHandler{*repo}
// }

// func (h *CartHandler) CreateTransaction(c *gin.Context) {
// 	// bind incoming http request
// 	request := model.CreateTransactionRequest{}
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		response.FailOrError(c, http.StatusUnprocessableEntity, "Create Cart failed", err)
// 		return
// 	}

// 	product, er := h.Repository.GetProductByID(uint(request.ProductID))

// 	if er != nil {
// 		response.FailOrError(c, http.StatusNotFound, "product not found", er)
// 		return
// 	}

// 	total := request.Quantity * product.Price

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
// 	cart := entity.Transaction{
// 		Product:  *product,
// 		Quantity: request.Quantity,
// 		Total:    float64(total),
// 		UserID:   uint(userIDf),
// 	}
// 	err := h.Repository.CreateTransaction(&cart)
// 	if err != nil {
// 		response.FailOrError(c, http.StatusInternalServerError, "Create post failed", err)
// 		return
// 	}
// 	fmt.Print("2")

// 	//success response
// 	response.Success(c, http.StatusCreated, "Post creation succeeded", request)
// }
