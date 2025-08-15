package service

import (
	"errors"
	"os"
	"time"

	"github.com/EDEN-NN/hydra-api/internal/apperrors"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	userService *UserService
}

type AuthValidate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateAuthService(service *UserService) *AuthService {
	return &AuthService{userService: service}
}

var secret = []byte(os.Getenv("API-KEY"))

func (service *AuthService) CreateToken(body *AuthValidate) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": body.Username,
		"password": body.Password,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *AuthService) VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, apperrors.NewConflictError("auth", err)
	}

	if !token.Valid {
		return nil, apperrors.NewError(apperrors.EINVALID, "invalid token", errors.New("token given is not valid"))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, apperrors.NewError(apperrors.EINVALID, "invalid claims", errors.New("could not parse claims"))
	}
	return claims, nil
}
