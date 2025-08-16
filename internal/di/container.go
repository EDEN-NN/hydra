package di

import (
	"context"

	"github.com/EDEN-NN/hydra-api/infra/config"
	"github.com/EDEN-NN/hydra-api/infra/database/mongodb"
	"github.com/EDEN-NN/hydra-api/infra/repository"
	"github.com/EDEN-NN/hydra-api/infra/web/handler"
	"github.com/EDEN-NN/hydra-api/infra/web/middleware"
	"github.com/EDEN-NN/hydra-api/infra/web/routes"
	"github.com/EDEN-NN/hydra-api/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	AppConfig   *config.AppConfig
	MongoDB     *mongo.Database
	UserHandler *handler.UserHandler
	AuthHandler *handler.AuthHandler
	Router      *gin.Engine
}

func NewContainer(ctx context.Context) (*Container, error) {
	appConfig := config.LoadConfig()

	db, err := mongodb.Connect()
	if err != nil {
		return nil, err
	}

	companyRepo := repository.CreateCompanyRepository(db)
	userRepo := repository.CreateUserRepository(db)
	userService := service.CreateUserService(userRepo)
	companyService := service.CreateCompanyService(companyRepo)
	authService := service.CreateAuthService(userService)
	userHandler := handler.CreateUserHandler(userService)
	companyHandler := handler.CreateCompanyHandler(companyService)
	authHandler := handler.CreateAuthHandler(authService, userService)

	router := gin.Default()
	router.POST("/login", authHandler.Login)
	router.POST("/user", userHandler.CreateUser)
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.HandleAuth(authService))
	routes.SetupUserRoutes(router, userHandler)
	routes.SetupAuthRoutes(router, authHandler)
	routes.SetupCompanyRoutes(router, companyHandler)

	return &Container{
		AppConfig:   appConfig,
		MongoDB:     db,
		UserHandler: userHandler,
		Router:      router,
	}, nil
}
