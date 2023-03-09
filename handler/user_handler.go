package handler

import (
	"ecoplant/model"
	"ecoplant/repository"
	"ecoplant/sdk/crypto"
	sdk_jwt "ecoplant/sdk/jwt"
	"ecoplant/sdk/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	Repository repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) userHandler {
	return userHandler{*repo}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user model.RegisterUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	result, err := h.Repository.CreateUser(user)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create user failed", err)
		return
	}
	response.Success(c, http.StatusCreated, "Success create user", result)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var request model.LoginUser
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailOrError(c, http.StatusBadRequest, "bad request", err)
		return
	}

	user, err := h.Repository.FindByUsernameOrEmail(request.UsernameOrEmail)
	if err != nil {
		response.FailOrError(c, http.StatusNotFound, " email or username not found", err)
		return
	}

	err = crypto.ValidateHash(request.Password, user.Password)
	if err != nil {
		msg := "wrong password"
		response.FailOrError(c, http.StatusBadRequest, msg, errors.New(msg))
		return
	}

	tokenJwt, err := sdk_jwt.GenerateToken(user)
	if err != nil {
		response.FailOrError(c, http.StatusInternalServerError, "create token failed", err)
		return
	}

	response.Success(c, http.StatusOK, "login success", gin.H{
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
