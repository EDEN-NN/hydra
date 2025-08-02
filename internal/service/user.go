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

	userOutput := service.MapUserEntityToOutput(user)

	return userOutput, nil
}

func (service *UserService) FindByID(id string) (*dto.UserOutput, error) {
	user, err := service.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	userOutput := service.MapUserEntityToOutput(user)

	return userOutput, nil
}

func (service *UserService) FindAll() ([]*dto.UserOutput, error) {
	users, err := service.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var usersOutput []*dto.UserOutput
	for _, user := range users {
		usersOutput = append(usersOutput, service.MapUserEntityToOutput(user))
	}

	return usersOutput, nil
}

func (service *UserService) MapUserEntityToOutput(entity *entity.User) *dto.UserOutput {
	userOutput := &dto.UserOutput{
		ID:        entity.ID.Hex(),
		Username:  entity.Username,
		Email:     entity.Email,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	return userOutput
}
