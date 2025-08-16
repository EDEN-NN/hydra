package entity

import (
	"errors"
	"time"

	"github.com/EDEN-NN/hydra-api/internal/apperrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BiometryStatus string
type Reputation string

const (
	ACCEPTED BiometryStatus = "ACCEPTED"
	DENIED   BiometryStatus = "DENIED"
	WAITING  BiometryStatus = "WAITING"

	NO_INDEX        Reputation = "NO_INDEX"
	NOT_RECOMMENDED Reputation = "NOT_RECOMMENDED"
	BAD             Reputation = "BAD"
	REGULAR         Reputation = "REGULAR"
	GOOD            Reputation = "GOOD"
	EXCELLENT       Reputation = "EXCELLENT"
)

type Company struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Name           string             `json:"name" bson:"name"`
	RegistryNumber string             `json:"registryNumber" bson:"registryNumber"`
	Biometry       BiometryStatus     `json:"biometry" bson:"biometry"`
	HasProduct     bool               `json:"hasProduct" bson:"hasProduct"`
	Reputation     Reputation         `json:"reputation" bson:"reputation"`
	UrlSite        string             `json:"urlSite" bson:"urlSite"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt" bson:"UpdatedAt"`
}

func CreateCompany(name, registryNumber, urlSite string, reputation Reputation) (*Company, error) {
	company := &Company{
		ID:             primitive.NewObjectID(),
		Name:           name,
		RegistryNumber: registryNumber,
		Biometry:       WAITING,
		HasProduct:     false,
		Reputation:     reputation,
		UrlSite:        urlSite,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := company.IsValid()
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (company *Company) IsValid() error {

	if len(company.Name) <= 5 {
		return apperrors.NewConflictError("name", errors.New("company name should have at least 6 characters"))
	}

	if len(company.RegistryNumber) != 14 {
		return apperrors.NewConflictError("registry number", errors.New("invalid registry number"))
	}

	return nil
}

func (company *Company) ChangeName(name string) (*Company, error) {
	company.Name = name
	company.UpdatedAt = time.Now()

	err := company.IsValid()
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (company *Company) ChangeRegistryNumber(registryNumber string) (*Company, error) {
	company.RegistryNumber = registryNumber
	company.UpdatedAt = time.Now()

	err := company.IsValid()

	if err != nil {
		return nil, err
	}

	return company, nil
}

func (company *Company) ActiveProduct() {
	company.HasProduct = true
	company.UpdatedAt = time.Now()
}

func (company *Company) UpdatedBiometryStatus(status BiometryStatus) {
	company.Biometry = status
	company.UpdatedAt = time.Now()
}
