package handler

import (
	"github.com/EDEN-NN/hydra-api/internal/service"
	"github.com/EDEN-NN/hydra-api/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func CreateUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var input = &dto.CreateUserInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters: " + err.Error()})
		return
	}

	result, err := handler.service.Create(input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (handler *UserHandler) FindByUsername(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
	}

	username := body.Username

	result, err := handler.service.FindByUsername(username)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}
