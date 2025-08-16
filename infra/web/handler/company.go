package handler

import (
	"net/http"

	"github.com/EDEN-NN/hydra-api/internal/service"
	"github.com/EDEN-NN/hydra-api/pkg/dto"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *service.CompanyService
}

func CreateCompanyHandler(service *service.CompanyService) *CompanyHandler {
	return &CompanyHandler{service: service}
}

func (handler *CompanyHandler) FindAll(c *gin.Context) {
	result, err := handler.service.ListCompanies()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (handler *CompanyHandler) CreateCompany(c *gin.Context) {
	var input = &dto.CreateCompanyInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters: " + err.Error()})
		return
	}

	result, err := handler.service.CreateCompany(input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result})
}

func (handler *CompanyHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	result, err := handler.service.GetCompanyByID(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (handler *CompanyHandler) UpdateName(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
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

	err := handler.service.ChangeName(id, body.Name)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "name updated successfully"})
}

func (handler *CompanyHandler) UpdateRegistryNumber(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	var body struct {
		RegistryNumber string `json:"registryNumber"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	err := handler.service.ChangeName(id, body.RegistryNumber)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "cnpj updated successfully"})
}

func (handler *CompanyHandler) ActiveProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	err := handler.service.ActiveProduct(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "product actived", "company": id})
}

func (handler *CompanyHandler) DeactiveProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	err := handler.service.DeactiveProduct(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "product deactived", "company": id})
}
