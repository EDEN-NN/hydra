package service

import (
	"errors"

	"github.com/EDEN-NN/hydra-api/infra/repository"
	"github.com/EDEN-NN/hydra-api/internal/apperrors"
	"github.com/EDEN-NN/hydra-api/internal/domain/entity"
	"github.com/EDEN-NN/hydra-api/pkg/dto"
)

type CompanyService struct {
	Repository *repository.CompanyRepository
}

func CreateCompanyService(repository *repository.CompanyRepository) *CompanyService {
	return &CompanyService{Repository: repository}
}

func (service *CompanyService) CreateCompany(data *dto.CreateCompanyInput) (string, error) {
	company, err := entity.CreateCompany(data.Name, data.RegistryNumber, data.UrlSite, entity.Reputation(data.Reputation))
	if err != nil {
		return "", err
	}
	return service.Repository.Create(company)
}

func (service *CompanyService) GetCompanyByID(id string) (*dto.CompanyOutput, error) {
	company, err := service.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	companyOutput := service.MapEntityToDTO(company)
	return companyOutput, nil
}

func (service *CompanyService) ChangeName(id, name string) error {
	company, err := service.Repository.FindByID(id)
	if err != nil {
		return err
	}
	err = company.ChangeName(name)
	if err != nil {
		return err
	}

	return service.Repository.UpdateCompany(company)
}

func (service *CompanyService) ListCompanies() ([]*dto.CompanyOutput, error) {
	list, err := service.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	companyOutputList := []*dto.CompanyOutput{}
	for _, company := range list {
		companyOutputList = append(companyOutputList, service.MapEntityToDTO(company))
	}

	return companyOutputList, err
}

func (service *CompanyService) UpdateRegistryNumber(id, registryNumber string) error {
	company, err := service.Repository.FindByID(id)
	if err != nil {
		return err
	}

	err = company.ChangeRegistryNumber(registryNumber)
	if err != nil {
		return err
	}

	return service.Repository.UpdateCompany(company)
}

func (service *CompanyService) ActiveProduct(id string) error {
	company, err := service.Repository.FindByID(id)
	if err != nil {
		return err
	}

	if company.HasProduct {
		return apperrors.NewError(apperrors.EINVALID, "unable to active product", errors.New("company already have product actived"))
	}

	company.ActiveProduct()
	return service.Repository.UpdateCompany(company)
}

func (service *CompanyService) DeactiveProduct(id string) error {
	company, err := service.Repository.FindByID(id)
	if err != nil {
		return err
	}

	if !company.HasProduct {
		return apperrors.NewError(apperrors.EINVALID, "unable to active product", errors.New("company already have product actived"))
	}

	company.DeactiveProduct()
	return service.Repository.UpdateCompany(company)
}

func (service *CompanyService) MapEntityToDTO(company *entity.Company) *dto.CompanyOutput {
	return &dto.CompanyOutput{
		ID:             company.ID.Hex(),
		Name:           company.Name,
		RegistryNumber: company.RegistryNumber,
		Biometry:       string(company.Biometry),
		HasProduct:     company.HasProduct,
		Reputation:     string(company.Reputation),
		UrlSite:        company.UrlSite,
		CreatedAt:      company.CreatedAt,
		UpdatedAt:      company.UpdatedAt,
	}
}
