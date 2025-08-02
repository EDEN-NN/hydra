package service

import (
	"context"
	"github.com/EDEN-NN/hydra-api/infra/repository"
	"github.com/EDEN-NN/hydra-api/internal/domain/entity"
	"github.com/EDEN-NN/hydra-api/pkg/dto"
)

type UserService struct {
	Repository *repository.UserRepository
}

func CreateUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}

func (service *UserService) Create(ctx context.Context, data *dto.CreateUserInput) (*string, error) {
	user, err := entity.CreateUser(
		data.Username,
		data.Password,
		data.Email,
		data.Name,
	)

	if err != nil {
		return nil, err
	}

	result, err := service.Repository.Create(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
