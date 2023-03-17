package handler

import (
	"ecoplant/entity"
	"ecoplant/repository"
	"ecoplant/sdk/crypto"
	sdk_jwt "ecoplant/sdk/jwt"
	"ecoplant/sdk/response"
	"errors"
	"net/http"
	"strconv"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type userHandler struct {
	Repository repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) userHandler {
	return userHandler{*repo}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user entity.RegisterUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	result, err := h.Repository.CreateUser(user)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "Username atau email sudah digunakan", err)
		return
	}

	response.Success(c, http.StatusCreated, "Success create user", result)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var request entity.LoginUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	user, err := h.Repository.FindByUsernameOrEmail(request.UsernameOrEmail)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, "Username atau email tidak ditemukan", err)
		return
	}

	err = crypto.ValidateHash(request.Password, user.Password)
	if err != nil {
		msg := "kata sandi salah, silakan coba lagi"
		response.FailOrError(c, http.StatusBadRequest, msg, errors.New(msg))
		return
	}

	tokenJwt, err := sdk_jwt.GenerateToken(user)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create token failed", err)
		return
	}

	response.Success(c, http.StatusOK, "Berhasil masuk", gin.H{
		"token": tokenJwt,
	})
}

func (h *userHandler) GetUserById(c *gin.Context) {
	id := c.Query("id")
	parsedID, _ := strconv.ParseUint(id, 10, 64)
	result, err := h.Repository.GetUserById(uint(parsedID))
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "get user failed", err)
		return
	}
	response.Success(c, http.StatusOK, "success get user", result)
}

func (h *userHandler) GetUserByBearer(c *gin.Context) {
	result, exist := c.Get("user")
	if !exist {
		response.FailOrError(c, http.StatusUnprocessableEntity, "no user key found", errors.New("Error"))
		return
	}
	claims, ok := result.(jwt.MapClaims)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error parsing ", errors.New("Error"))
		return
	}

	userIDc := claims["id"]
	userIDf, ok := userIDc.(float64)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error get id", errors.New("Error"))
		return
	}

	user, err := h.Repository.GetUserById(uint(userIDf))
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "Failed get User", err)
	}

	response.Success(c, http.StatusCreated, "Get user Success", user)
}

func (h *userHandler) UpdateUser(c *gin.Context) {

	var request entity.UpdateUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	result, exist := c.Get("user")
	if !exist {
		response.FailOrError(c, http.StatusUnprocessableEntity, "no user key found", errors.New("Error"))
		return
	}
	claims, ok := result.(jwt.MapClaims)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error parsing ", errors.New("Error"))
		return
	}

	userIDc := claims["id"]
	userIDf, ok := userIDc.(float64)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error get id", errors.New("Error"))
		return
	}

	var userUpdate entity.User = entity.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
	}

	err = h.Repository.UpdateUser(uint(userIDf), userUpdate)
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error updating user", err)
		return
	}

	response.Success(c, http.StatusOK, "Success updating profile", request)
}

func (h *userHandler) UpdateProfilePicture(c *gin.Context) {
	supClient := supabasestorageuploader.NewSupabaseClient(
		"https://oybixjqqpdbzadyzeeml.supabase.co",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im95Yml4anFxcGRiemFkeXplZW1sIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NzgwNzkwOTYsImV4cCI6MTk5MzY1NTA5Nn0.uxwKBWc9kl4IOxWJMrKUHxDnJbQ19JNgJfbo3oJYiAI",
		"ecoplants",
		"profile",
	)

	result, exist := c.Get("user")
	if !exist {
		response.FailOrError(c, http.StatusUnprocessableEntity, "no user key found", errors.New("Error"))
		return
	}
	claims, ok := result.(jwt.MapClaims)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error parsing ", errors.New("Error"))
		return
	}

	userIDc := claims["id"]
	userIDf, ok := userIDc.(float64)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error get id", errors.New("Error"))
		return
	}

	file, err := c.FormFile("picture")
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error getting request", err)
		return
	}
	link, err := supClient.Upload(file)
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Failed uploading picture", err)
		return
	}

	err = h.Repository.UpdateProfilePicture(uint(userIDf), link)
	if err != nil {
		response.FailOrError(c, http.StatusUnprocessableEntity, "Failed updating picture", err)
		return
	}

	response.Success(c, http.StatusOK, "Success updating profile", link)

}
