package routes

import (
	"github.com/EDEN-NN/hydra-api/infra/web/handler"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, handler *handler.AuthHandler) {
	authRoutes := router.Group("/api/login")
	{
		authRoutes.POST("/", handler.Login)
	}
}
