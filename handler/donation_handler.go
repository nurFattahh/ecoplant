package handler

import (
	"ecoplant/entity"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"errors"
	"net/http"
	"strconv"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	supClient := supabasestorageuploader.NewSupabaseClient(
		"https://oybixjqqpdbzadyzeeml.supabase.co",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im95Yml4anFxcGRiemFkeXplZW1sIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NzgwNzkwOTYsImV4cCI6MTk5MzY1NTA5Nn0.uxwKBWc9kl4IOxWJMrKUHxDnJbQ19JNgJfbo3oJYiAI",
		"ecoplants",
		"donasi",
	)

	filePicture, err := c.FormFile("picture")
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error getting picture", err)
		return
	}
	linkPicture, err := supClient.Upload(filePicture)
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Failed uploading picture", err)
		return
	}

	name := c.PostForm("name")
	regency := c.PostForm("regency")
	district := c.PostForm("district")
	target := c.PostForm("target")
	parseTarget, _ := strconv.Atoi(target)
	remainDay := c.PostForm("remain_day")
	parseDay, _ := strconv.Atoi(remainDay)
	plan := c.PostForm("plan")
	news := c.PostForm("news")

	donation := entity.Donation{
		Name:        name,
		Regency:     regency,
		District:    district,
		Target:      float64(parseTarget),
		Picture:     linkPicture,
		RemainDay:   parseDay,
		Plan:        plan,
		News:        news,
		Community:   *community,
		CommunityID: uint(parsedID),
	}

	err = h.Repository.CreateDonation(&donation)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Create donation failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "donation creation succeeded", donation)
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
		"donation":   donation,
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

//USER DONATION

func (h *DonationHandler) UserDonation(c *gin.Context) {
	id := c.Param("id")

	parsedID, _ := strconv.ParseUint(id, 10, 64)

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

	request := entity.UserDonationRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Create transaction failed", err)
		return
	}

	donation, err := h.Repository.GetDonationByID(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "failed getting donation", err)
		return
	}

	donation.Wallet += request.Nominal
	donation.NumDonate++

	nominal := request.Nominal
	switch nominal {
	case 1:
		nominal = 5000
	case 2:
		nominal = 10000
	case 3:
		nominal = 20000
	case 4:
		nominal = 50000
	case 5:
		nominal = 100000
	case 6:
		nominal = 1000000
	default:
		nominal = request.Nominal
	}

	payMethodReq := request.PaymentMethod
	var payMethod string
	switch payMethodReq {
	case 1:
		payMethod = "Bank BCA"
	case 2:
		payMethod = "Bank BRI"
	case 3:
		payMethod = "Bank BNI"
	case 4:
		payMethod = "Bank Mandiri"
	case 5:
		payMethod = "Bank CIMBNIAGA"
	}

	var donate entity.UserDonation = entity.UserDonation{
		UserID:        uint(userIDf),
		DonationID:    uint(parsedID),
		Nominal:       nominal,
		PaymentMethod: payMethod,
		Donation:      *donation,
	}

	err = h.Repository.CreateUserDonation(float64(parsedID), nominal, &donate)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "failed create donation", err)
		return
	}

	response.Success(c, http.StatusOK, "Donation Success", donate)

}

func (h *DonationHandler) GetAllUserDonation(c *gin.Context) {
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

	donations, err := h.Repository.GetAllUserDonation(uint(userIDf))
	if err != nil {
		response.FailOrError(c, http.StatusOK, "failed get all donation", err)
		return
	}

	response.Success(c, http.StatusOK, "Success getting user donations", donations)
}

func (h *DonationHandler) UpdatePlanAndNewsDonation(c *gin.Context) {
	id := c.Param("id")

	parsedID, _ := strconv.ParseUint(id, 10, 64)

	request := entity.UpdatePlanAndNewsDonation{}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "updating failed", err)
		return
	}

	donate := entity.Donation{
		Plan: request.Plan,
		News: request.News,
	}

	err := h.Repository.UpdatePlanAndNewsDonation(uint(parsedID), donate)
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "failed updating plan and news", err)
		return
	}

	response.Success(c, http.StatusOK, "Success updating plan and news", request)

}
