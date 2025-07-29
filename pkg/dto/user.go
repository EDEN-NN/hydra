package dto

import "time"

type CreateUserInput struct {
	Username string `json:"username" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=3"`
}

type CreateUserOutput struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
