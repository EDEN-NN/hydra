package handler

import (
	"github.com/EDEN-NN/hydra-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
	userService *service.UserService
}

func CreateAuthHandler(authService *service.AuthService, userService *service.UserService) *AuthHandler {
	return &AuthHandler{authService: authService, userService: userService}
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var body = &service.AuthValidate{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params to login"})
		return
	}

	err := handler.userService.ValidateLogin(body.Username, body.Password)
	if err != nil {
		c.Error(err)
		return
	}

	tokenString, err := handler.authService.CreateToken(body)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
