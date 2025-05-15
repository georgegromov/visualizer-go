package service

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
	"visualizer-go/internal/repository"

	"github.com/google/uuid"
)

type UserService struct {
	log  *slog.Logger
	repo repository.User
}

func NewUserService(log *slog.Logger, repo repository.User) *UserService {
	return &UserService{
		log:  log,
		repo: repo,
	}
}

func (us *UserService) Login(ctx context.Context, dto dto.UserLoginDto) (models.User, string, error) {
	const op = "service.UserService.Login"

	// TODO:
	// hash compare passwords

	user, err := us.GetByUsername(ctx, dto.Username)
	if err != nil {
		us.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.User{}, "", fmt.Errorf("%s: %w", op, repository.ErrInvalidCredentials)
	}

	trimmedPasswordHash := strings.ReplaceAll(user.PasswordHash, " ", "")

	if trimmedPasswordHash != dto.Password {
		return models.User{}, "", fmt.Errorf("%s: %w", op, repository.ErrInvalidCredentials)
	}

	return user, "test_token_123", nil
}

func (us *UserService) GetByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	// const op = "service.UserService.GetByID"
	return us.repo.GetByID(ctx, userID)
}

func (us *UserService) GetByUsername(ctx context.Context, username string) (models.User, error) {
	// const op = "service.UserService.GetByUsername"
	return us.repo.GetByUsername(ctx, username)
}

func (us *UserService) Create(ctx context.Context, dto dto.UserCreateDto) error {
	// const op = "service.UserService.Create"
	return us.repo.Create(ctx, dto)
}

func (us *UserService) Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error {
	// const op = "service.UserService.Update"
	return us.repo.Update(ctx, userID, dto)
}
