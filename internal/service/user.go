package service

import (
	"context"
	"fmt"
	"log/slog"
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

func (us *UserService) Login(ctx context.Context, dto dto.UserLoginDto) (models.UserWithToken, error) {
	const op = "service.UserService.Login"

	// 1. Find user by username +
	// 2. Compare passwords +
	// 3. Generate token -
	// 4. Return user with token +

	foundUser, err := us.GetByUsername(ctx, dto.Username)
	if err != nil {
		us.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.UserWithToken{}, err
	}

	if err = foundUser.ComparePasswords(dto.Password); err != nil {
		us.log.Error(fmt.Sprintf("%s: %v", op, err))
		return models.UserWithToken{}, fmt.Errorf("%s: %w", op, repository.ErrInvalidCredentials)
	}

	// !!! IMPORTANT DO NOT REMOVE !!! This remove password_hash from json struct
	foundUser.SanitizePassword()

	return models.UserWithToken{
		User:  foundUser,
		Token: "test_token_123",
	}, nil
}

func (us *UserService) GetByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	// const op = "service.UserService.GetByID"
	return us.repo.GetByID(ctx, userID)
}

func (us *UserService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	// const op = "service.UserService.GetByUsername"
	return us.repo.GetByUsername(ctx, username)
}

func (us *UserService) Create(ctx context.Context, user *models.User) error {
	// const op = "service.UserService.Create"

	// 1. Find user by username
	// 2. Prepare create (hash password)
	// 3. Create user

	foundUser, err := us.GetByUsername(ctx, user.Username)

	if foundUser != nil || err == nil {
		fmt.Print("found user:", foundUser.Username)
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	if err := user.PrepareCreate(); err != nil {
		us.log.Error("error occured while prepairing user to create")
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	return us.repo.Create(ctx, user)
}

func (us *UserService) Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error {
	// const op = "service.UserService.Update"
	return us.repo.Update(ctx, userID, dto)
}
