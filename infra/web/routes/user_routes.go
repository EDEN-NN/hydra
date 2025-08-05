package routes

import (
	"github.com/EDEN-NN/hydra-api/infra/web/handler"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	userRoutes := router.Group("/api/users")
	{
		userRoutes.GET("/", userHandler.FindAll)
		userRoutes.GET("/user", userHandler.FindByUsername)
		userRoutes.GET("/email", userHandler.FindByEmail)
		userRoutes.GET("/user/:id", userHandler.FindByID)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.POST("/user/change-name/:id", userHandler.UpdateName)
		userRoutes.POST("/user/change-email/:id", userHandler.ChangeEmail)
	}
}
