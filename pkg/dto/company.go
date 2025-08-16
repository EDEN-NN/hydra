package dto

import "time"

type CreateCompanyInput struct {
	Name           string `json:"name" binding:"required,min=5"`
	RegistryNumber string `json:"registryNumber" binding:"required,min=14,max=14"`
	UrlSite        string `json:"urlSite" binding:"required,url"`
	Reputation     string `json:"reputation" binding:"required"`
}

type CompanyOutput struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	RegistryNumber string    `json:"registryNumber"`
	Biometry       string    `json:"biometry"`
	HasProduct     bool      `json:"hasProduct"`
	Reputation     string    `json:"reputation"`
	UrlSite        string    `json:"urlSite"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
