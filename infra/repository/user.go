package repository

import (
	"context"
	"errors"
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
	var userEntity = &entity.User{}
	userWithSameUsernameOrEmail, err := repository.DB.Collection("users").Find(context.Background(), bson.D{{"username", data.Username}})

	_ = userWithSameUsernameOrEmail.Decode(userEntity)

	if userWithSameUsernameOrEmail != nil {
		err := apperrors.NewError(apperrors.EINVALID, "username or email already in use", errors.New("invalid params"))
		err.Metadata = map[string]interface{}{
			"invalid data": []string{data.Email, data.Username},
		}
		return nil, err
	}

	result, err := repository.DB.Collection("users").InsertOne(context.Background(), &data)
	if err != nil {
		return nil, apperrors.NewError(apperrors.EINVALID, "fail to comunicate with database", err)
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("new document inserted: %s", result.InsertedID)

	return &userID, nil
}
