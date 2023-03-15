package handler

import (
	"ecoplant/entity"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DonationHandler struct {
	Repository repository.DonationRepository
}

func NewDonationHandler(repo *repository.DonationRepository) DonationHandler {
	return DonationHandler{*repo}
}

func (h *DonationHandler) CreateDonation(c *gin.Context) {
	id := c.Param("id")
	parsedID, _ := strconv.ParseUint(id, 10, 64)

	community, err := h.Repository.GetCommunityByID(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Failed getting community", err)
		return
	}

	request := entity.CreateDonation{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "bad request", err)
		return
	}

	donation := entity.Donation{
		Name:        request.Name,
		Regency:     request.Regency,
		District:    request.District,
		Target:      request.Target,
		RemainDay:   request.RemainDay,
		Plan:        request.Plan,
		News:        request.News,
		Community:   *community,
		CommunityID: uint(parsedID),
	}

	err = h.Repository.CreateDonation(&donation)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Create donation failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "donation creation succeeded", request)
}

func (h *DonationHandler) GetAllDonation(c *gin.Context) {
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
	donation, totalElements, err := h.Repository.GetAllDonation(&productParam)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Donation not found", err)
		return
	}

	productParam.ProcessPagin(totalElements)
	response.Success(c, http.StatusOK, "Donation Found", gin.H{
		"pagination": &productParam,
		"community":  donation,
	})
}

func (h *DonationHandler) GetDonationByID(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid id params", err)
		return
	}

	donation, err := h.Repository.GetDonationByID(uint(parsedID))

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "donation not found", err)
		return
	}

	response.Success(c, http.StatusOK, "community found", donation)
}

func (h *DonationHandler) GetDonationByRegency(c *gin.Context) {
	query := c.Query("regency")

	products, err := h.Repository.GetDonationByRegency(query)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid regency community", err)
		return
	}

	response.Success(c, http.StatusOK, "donation found", products)
}
