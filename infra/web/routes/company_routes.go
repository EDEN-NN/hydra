package routes

import (
	"github.com/EDEN-NN/hydra-api/infra/web/handler"
	"github.com/gin-gonic/gin"
)

func SetupCompanyRoutes(router *gin.Engine, handler *handler.CompanyHandler) {
	companyRoutes := router.Group("/api/companies")
	{
		companyRoutes.GET("/", handler.FindAll)
		companyRoutes.GET("/company/:id", handler.FindByID)
		companyRoutes.POST("/company", handler.CreateCompany)
		companyRoutes.PATCH("/company/:id/active", handler.ActiveProduct)
		companyRoutes.PATCH("/company/:id/deactive", handler.DeactiveProduct)
		companyRoutes.PATCH("/company/:id/change-name", handler.UpdateName)
		companyRoutes.PATCH("/company/:id/change-registry-number", handler.UpdateRegistryNumber)
	}
}
