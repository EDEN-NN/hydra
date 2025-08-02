package repository

import (
	"context"
	"github.com/EDEN-NN/hydra-api/internal/apperrors"
	"github.com/EDEN-NN/hydra-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepository struct {
	DB *mongo.Database
}

func CreateUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) Create(data *entity.User) (*string, error) {
	result, err := repository.DB.Collection("users").InsertOne(context.Background(), &data)
	if err != nil {
		return nil, apperrors.NewError(apperrors.EINVALID, "fail to insert a new user", err)
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("new document inserted: %s", result.InsertedID)

	return &userID, nil
}

func (repository *UserRepository) FindByUsername(username string) (*entity.User, error) {
	var userEntity = &entity.User{}
	err := repository.DB.Collection("users").FindOne(context.Background(), bson.D{{"username", username}}).Decode(userEntity)
	if err != nil {
		return nil, apperrors.NewError(apperrors.EINVALID, "error searching for user", err)
	}

	return userEntity, nil
}
