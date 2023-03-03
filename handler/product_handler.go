package handler

import (
	"ecoplant/entity"
	"ecoplant/model"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"net/http"
	"strconv"

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
	err := h.Repository.CreateProduct(&product)
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

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid id params", err)
		return
	}

	product, err := h.Repository.GetProductByID(uint(parsedID))

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "product not found", err)
		return
	}

	response.Success(c, http.StatusOK, "product found", product)
}

func (h *ProductHandler) GetProductByName(c *gin.Context) {
	query := c.Param("name")

	product, err := h.Repository.GetProductByName(query)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid name product", err)
		return
	}

	response.Success(c, http.StatusOK, "product found", product)
}

// func (h *ProductHandler) GetListProduct(ctx *gin.Context) {
// 	var postParam model.PostParam

// 	if err := h.BindParam(ctx, &postParam); err != nil {
// 		h.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", nil)
// 		return
// 	}

// 	postParam.FormatPagination()

// 	var posts []entity.Post

// 	if err := h.db.
// 		Model(entity.Product{}).
// 		Limit(int(postParam.Limit)).
// 		Offset(int(postParam.Offset)).
// 		Find(&posts).Error; err != nil {
// 		response.FailOrError(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}

// 	var totalElements int64

// 	if err := h.db.
// 		Model(entity.Post{}).
// 		Limit(int(postParam.Limit)).
// 		Offset(int(postParam.Offset)).
// 		Count(&totalElements).Error; err != nil {
// 		h.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}

// 	postParam.ProcessPagination(totalElements)

// 	h.Success(ctx, http.StatusOK, "Successfully get list post", posts, &postParam.PaginationParam)
// }
