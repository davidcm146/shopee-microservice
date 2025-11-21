package service

import (
	"context"
	"errors"
	"time"

	"github.com/davidcm146/shopee-microservice/user-service/internal/dto"
	"github.com/davidcm146/shopee-microservice/user-service/internal/models"
	"github.com/davidcm146/shopee-microservice/user-service/internal/repository"
	"github.com/davidcm146/shopee-microservice/user-service/pkg/db"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, input *dto.RegisterInput) (*models.User, error)
	Login(ctx context.Context, input *dto.LoginInput) (*dto.LoginResult, error)
	Logout(ctx context.Context, sessionID string) error
	GetUserBySessionID(ctx context.Context, sessionID string) (*models.User, error)
}

type authService struct {
	userRepo repository.UserRepository
	redis    *db.RedisConfig
}

func NewAuthService(userRepo repository.UserRepository, redis *db.RedisConfig) AuthService {
	return &authService{
		userRepo: userRepo,
		redis:    redis,
	}
}

// var validate = validator.New()

func (s *authService) Register(ctx context.Context, input *dto.RegisterInput) (*models.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(ctx, input.Email)
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	// Create user
	user, err := models.NewUser(input.Email, input.Password, input.Name, input.Role)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *authService) Login(ctx context.Context, input *dto.LoginInput) (*dto.LoginResult, error) {
	user, err := s.userRepo.FindByEmail(ctx, input.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	sessionID := uuid.NewString()
	err = s.redis.Client.Set(ctx, sessionID, user.ID, time.Hour*24).Err()
	if err != nil {
		return nil, err
	}

	return &dto.LoginResult{
		SessionID: sessionID,
	}, nil
}

func (s *authService) Logout(ctx context.Context, sessionID string) error {
	return s.redis.Client.Del(ctx, sessionID).Err()
}

func (s *authService) GetUserBySessionID(ctx context.Context, sessionID string) (*models.User, error) {
	userIDStr, err := s.redis.Client.Get(ctx, sessionID).Result()
	if err != nil {
		return nil, errors.New("session not found")
	}

	return s.userRepo.FindByID(ctx, userIDStr)
}
