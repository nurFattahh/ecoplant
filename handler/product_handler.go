package handler

import (
	"ecoplant/entity"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Repository repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) ProductHandler {
	return ProductHandler{*repo}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	request := entity.CreateProduct{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Create product failed", err)
		return
	}

	product := entity.Product{
		Name:        request.Name,
		Price:       request.Price,
		Rating:      request.Rating,
		Description: request.Description,
		Merchant:    request.Merchant,
		Picture:     request.Picture,
		Regency:     request.Regency,
		District:    request.District,
	}
	err := h.Repository.CreateProduct(&product)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Create product failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "Product creation succeeded", request)
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	queryLimit := c.Query("limit")
	queryPage := c.Query("page")

	parseLimit, _ := strconv.ParseInt(queryLimit, 10, 64)
	parsePage, _ := strconv.ParseInt(queryPage, 10, 64)

	var productParam entity.PaginParam
	productParam.Limit = int(parseLimit)
	productParam.Page = int(parsePage)
	if err := h.Repository.BindParam(c, &productParam); err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid request body", err)
		return
	}
	productParam.FormatPagin()
	products, totalElements, err := h.Repository.GetAllProduct(&productParam)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Product not found", err)
		return
	}

	productParam.ProcessPagin(totalElements)
	response.Success(c, http.StatusOK, "Product Found", gin.H{
		"pagination": &productParam,
		"products":   products,
	})
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
	query := c.Query("name")

	products, err := h.Repository.GetProductByName(query)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid name product", err)
		return
	}

	response.Success(c, http.StatusOK, "product found", products)
}

func (h *ProductHandler) DeleteProductById(c *gin.Context) {
	ID := c.Param("id")

	parsedID, _ := strconv.ParseUint(ID, 10, 64)

	err := h.Repository.DeleteProduct(uint(parsedID))

	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "delete product failed", err)
		return
	}

	response.Success(c, http.StatusOK, "successfully deleted product", nil)
}

func (h *ProductHandler) UpdateLocation(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid id params", err)
		return
	}
	request := entity.UpdateLocation{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Failed get request", err)
		return
	}

	err = h.Repository.UpdateProduct(uint(parsedID), request.Location)
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Update product failed", err)
		return
	}

	response.Success(c, http.StatusOK, "successfully update product", request)

}
