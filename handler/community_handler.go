package handler

import (
	"ecoplant/entity"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"net/http"
	"strconv"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
)

type CommunityHandler struct {
	Repository repository.CommunityRepository
}

func NewCommunityHandler(repo *repository.CommunityRepository) CommunityHandler {
	return CommunityHandler{*repo}
}

func (h *CommunityHandler) CreateCommunity(c *gin.Context) {
	supClient := supabasestorageuploader.NewSupabaseClient(
		"https://oybixjqqpdbzadyzeeml.supabase.co",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im95Yml4anFxcGRiemFkeXplZW1sIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NzgwNzkwOTYsImV4cCI6MTk5MzY1NTA5Nn0.uxwKBWc9kl4IOxWJMrKUHxDnJbQ19JNgJfbo3oJYiAI",
		"ecoplants",
		"komunitas",
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
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	description := c.PostForm("description")

	fileDocument, err := c.FormFile("document")
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error getting document", err)
		return
	}
	linkDocument, err := supClient.Upload(fileDocument)
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Failed uploading picture", err)
		return
	}

	community := entity.Community{
		Picture:     linkPicture,
		Name:        name,
		Email:       email,
		Phone:       phone,
		Description: description,
		Document:    linkDocument,
	}
	err = h.Repository.CreateCommunity(&community)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Create Community failed", err)
		return
	}

	response.Success(c, http.StatusCreated, "Community creation succeeded", community)
}

func (h *CommunityHandler) GetAllCommunity(c *gin.Context) {
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
	community, totalElements, err := h.Repository.GetAllCommunity(&productParam)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Community not found", err)
		return
	}

	productParam.ProcessPagin(totalElements)
	response.Success(c, http.StatusOK, "Community Found", gin.H{
		"pagination": &productParam,
		"community":  community,
	})
}

func (h *CommunityHandler) GetCommunityByID(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid id params", err)
		return
	}

	community, err := h.Repository.GetCommunityByID(uint(parsedID))

	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "community not found", err)
		return
	}

	response.Success(c, http.StatusOK, "community found", community)
}

func (h *CommunityHandler) GetCommunityByName(c *gin.Context) {
	query := c.Query("name")

	products, err := h.Repository.GetCommunityByName(query)

	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "invalid name community", err)
		return
	}

	response.Success(c, http.StatusOK, "community found", products)
}
