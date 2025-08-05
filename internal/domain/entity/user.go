package entity

import (
	"errors"
	"github.com/EDEN-NN/hydra-api/internal/apperrors"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"time"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

func (user *User) IsValid() error {

	if len(user.Username) < 6 {
		return apperrors.NewConflictError("username", errors.New("your username should have at least 6 characters"))
	}

	if len(user.Password) < 8 {
		return apperrors.NewConflictError("password", errors.New("your password should have at least 8 characters"))
	}

	if len(user.Name) < 3 {
		return apperrors.NewConflictError("name", errors.New("your name should have at least 3 characters"))
	}

	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		apperrors.NewConflictError("email", errors.New("invalid email format"))
	}
	return nil
}

func CreateUser(username, password, email, name string) (*User, error) {
	newUser := &User{
		ID:        primitive.NewObjectID(),
		Username:  username,
		Password:  password,
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := newUser.IsValid()

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (user *User) ChangeName(name string) error {
	user.Name = name
	user.UpdatedAt = time.Now()

	err := user.IsValid()

	if err != nil {
		return err
	}

	return nil
}

func (user *User) ChangeEmail(email string) error {
	user.Email = email
	user.UpdatedAt = time.Now()

	err := user.IsValid()
	if err != nil {
		return err
	}

	return nil
}

func GenerateHashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", apperrors.NewError(apperrors.EINTERNAL, "error while creating or updating user", errors.New("fail trying to create a hashed password"))
	}

	return string(hashedBytes), nil
}

func CompareHash(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return apperrors.NewConflictError("user", apperrors.NewError(apperrors.ECONFLICT, "username or password invalid", err))
	}
	return nil
}
