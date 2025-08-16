package handler

import (
	"net/http"

	"github.com/EDEN-NN/hydra-api/internal/service"
	"github.com/EDEN-NN/hydra-api/pkg/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func CreateUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) FindAll(c *gin.Context) {
	result, err := handler.service.FindAll()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
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

func (handler *UserHandler) FindByID(c *gin.Context) {
	idFromParams := c.Param("id")
	if &idFromParams == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	result, err := handler.service.FindByID(idFromParams)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (handler *UserHandler) FindByEmail(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	result, err := handler.service.FindByEmail(body.Email)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (handler *UserHandler) UpdateName(c *gin.Context) {
	idFromRequest := c.Param("id")
	if idFromRequest == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	var body struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	err := handler.service.ChangeName(body.Name, idFromRequest)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "name updated successfully"})
}

func (handler *UserHandler) ChangeEmail(c *gin.Context) {
	var idFromParam = c.Param("id")
	if idFromParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	var body struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	err := handler.service.ChangeEmail(body.Email, idFromParam)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "email changed successfully"})

}
