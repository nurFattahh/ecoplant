package handler

import (
	"ecoplant/model"
	"ecoplant/repository"
	"ecoplant/sdk/response"
	"net/http"

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
	response.Success(c, http.StatusInternalServerError, "Success create user", result)
}
