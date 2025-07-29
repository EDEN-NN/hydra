package di

import (
	"context"
	"github.com/EDEN-NN/hydra-api/infra/config"
	"github.com/EDEN-NN/hydra-api/infra/database/mongodb"
	"github.com/EDEN-NN/hydra-api/infra/repository"
	"github.com/EDEN-NN/hydra-api/infra/web/handler"
	"github.com/EDEN-NN/hydra-api/infra/web/routes"
	"github.com/EDEN-NN/hydra-api/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	AppConfig   *config.AppConfig
	MongoDB     *mongo.Database
	UserHandler *handler.UserHandler
	Router      *gin.Engine
}

func NewContainer(ctx context.Context) (*Container, error) {
	appConfig := config.LoadConfig()

	db, err := mongodb.Connect()
	if err != nil {
		return nil, err
	}

	userRepo := repository.CreateUserRepository(db)
	userService := service.CreateUserService(userRepo)
	userHandler := handler.CreateUserHandler(userService)

	router := gin.Default()
	routes.SetupUserRoutes(router, userHandler)

	return &Container{
		AppConfig:   appConfig,
		MongoDB:     db,
		UserHandler: userHandler,
		Router:      router,
	}, nil
}
