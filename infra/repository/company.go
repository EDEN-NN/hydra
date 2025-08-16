package repository

import (
	"context"
	"log"

	"github.com/EDEN-NN/hydra-api/internal/apperrors"
	"github.com/EDEN-NN/hydra-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepository struct {
	DB *mongo.Database
}

func CreateCompanyRepository(db *mongo.Database) *CompanyRepository {
	return &CompanyRepository{
		DB: db,
	}
}

func (repository *CompanyRepository) Create(data *entity.Company) (string, error) {
	result, err := repository.DB.Collection("companies").InsertOne(context.Background(), &data)
	if err != nil {
		return "", apperrors.NewError(apperrors.EINVALID, "fail to insert a new company", err)
	}

	companyID := result.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("new company inserted: %s", result.InsertedID)

	return companyID, nil
}

func (repository *CompanyRepository) FindByID(id string) (*entity.Company, error) {
	var companyEntity = &entity.Company{}
	companyId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, apperrors.NewError(apperrors.EINVALID, "invalid id", err)
	}

	err = repository.DB.Collection("companies").
		FindOne(context.Background(), bson.D{{Key: "_id", Value: companyId}}).
		Decode(companyEntity)
	if err != nil {
		return nil, apperrors.NewError(apperrors.ENOTFOUND, "company not found", err)
	}

	return companyEntity, nil
}

func (repository *CompanyRepository) FindByName(name string) (*entity.Company, error) {
	var companyEntity = &entity.Company{}
	err := repository.DB.Collection("companies").
		FindOne(context.Background(), bson.D{{Key: "name", Value: name}}).
		Decode(companyEntity)
	if err != nil {
		return nil, apperrors.NewError(apperrors.ENOTFOUND, "company not found", err)
	}

	return companyEntity, nil
}

func (repository *CompanyRepository) UpdateCompany(company *entity.Company) error {
	filter := bson.M{"_id": company.ID}
	update := bson.M{"$set": &company}
	_, err := repository.DB.Collection("companies").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return apperrors.NewNotFoundError("company not found")
	}

	return nil
}

func (repository *CompanyRepository) FindAll() ([]*entity.Company, error) {
	var results []*entity.Company
	cursor, err := repository.DB.Collection("companies").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, apperrors.NewConflictError("companies", err)
	}

	err = cursor.All(context.Background(), &results)
	if err != nil {
		return nil, apperrors.NewConflictError("companies", err)
	}

	return results, nil
}
