package repository

import (
	"context"
	"github.com/EDEN-NN/hydra-api/internal/domain/entity"
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
		return nil, err
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("new document inserted: %s", result.InsertedID)

	return &userID, nil
}
