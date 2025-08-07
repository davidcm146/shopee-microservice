package dto

import "github.com/davidcm146/shopee-microservice/user-service/internal/models"

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	SessionID string       `json:"sessionID"`
	User      *models.User `json:"user"`
}
