package routes

import (
	"github.com/EDEN-NN/hydra-api/infra/web/handler"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	userRoutes := router.Group("/api/users")
	{
		userRoutes.GET("/user", userHandler.FindByUsername)
		userRoutes.POST("/", userHandler.CreateUser)
	}
}
