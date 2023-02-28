package handler

import (
	"ecoplant/entity"
	"ecoplant/model"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Repository repository.ProductRepository
}

func NewProductRepository(repo *repository.ProductRepository) ProductHandler {
	return ProductHandler{*repo}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	request := model.CreateProduct{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Create product failed", err)
		return
	}

	product := entity.Product{
		Name:        request.Name,
		Price:       request.Price,
		Rating:      request.Rating,
		Description: request.Description,
	}
	err := h.Repository.CreatePost(&product)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Create product failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "Post creation succeeded", request)
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	posts, err := h.Repository.GetAllProduct()
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "fail getting product", err)
		return
	}

	response.Success(c, http.StatusOK, "Success getting all product", posts)
}
