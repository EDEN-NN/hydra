package service

import (
	"github.com/EDEN-NN/hydra-api/infra/repository"
	"github.com/EDEN-NN/hydra-api/internal/domain/entity"
	"github.com/EDEN-NN/hydra-api/pkg/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserService struct {
	Repository *repository.UserRepository
}

func CreateUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}

func (service *UserService) ValidateLogin(username, password string) error {
	user, err := service.Repository.FindByUsername(username)
	if err != nil {
		return err
	}

	err = entity.CompareHash(password, user.Password)
	if err != nil {
		return err
	}

	return nil
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

func (service *UserService) FindByEmail(email string) (*dto.UserOutput, error) {
	user, err := service.Repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	userOutput := service.MapUserEntityToOutput(user)

	return userOutput, nil
}

func (service *UserService) ChangeName(name string, id string) error {
	userResult, err := service.Repository.FindByID(id)
	if err != nil {
		return err
	}

	err = userResult.ChangeName(name)
	if err != nil {
		return err
	}

	err = service.Repository.UpdateUser(userResult)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) ChangeEmail(email string, id string) error {
	userResult, err := service.Repository.FindByID(id)
	if err != nil {
		return err
	}

	err = userResult.ChangeEmail(email)
	if err != nil {
		return err
	}

	if err = service.Repository.UpdateUser(userResult); err != nil {
		return err
	}

	return nil
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

func (service *UserService) MapDtoOutputToEntity(output *dto.UserOutput) *entity.User {
	id, _ := primitive.ObjectIDFromHex(output.ID)
	userEntity := &entity.User{
		ID:        id,
		Username:  output.Username,
		Name:      output.Name,
		Email:     output.Email,
		UpdatedAt: time.Now(),
	}
	return userEntity
}
