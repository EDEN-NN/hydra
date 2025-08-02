package service

import (
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

func (service *UserService) Create(data *dto.CreateUserInput) (*string, error) {
	hashedPassword, _ := entity.GenerateHashPassword(data.Password)
	user, err := entity.CreateUser(
		data.Username,
		hashedPassword,
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

func (service *UserService) FindByUsername(username string) (*dto.UserOutput, error) {
	user, err := service.Repository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	userOutput := &dto.UserOutput{
		ID:        user.ID.Hex(),
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userOutput, nil
}
