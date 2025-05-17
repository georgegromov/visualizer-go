package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"visualizer-go/internal/domains/users"
	"visualizer-go/internal/domains/users/repository"
	jwt_manager "visualizer-go/pkg/jwt"

	"github.com/google/uuid"
)

type userUsecase struct {
	log        *slog.Logger
	repo       users.Repository
	jwtManager *jwt_manager.JwtManager
}

func NewUserService(log *slog.Logger, repo users.Repository, jwtManager *jwt_manager.JwtManager) users.Usecase {
	return &userUsecase{
		log:        log,
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (us *userUsecase) Login(ctx context.Context, dto users.UserLoginDto) (*users.UserWithToken, error) {
	const op = "usecase.userUsecase.Login"
	foundUser, err := us.GetByUsername(ctx, dto.Username)
	if err != nil {
		us.log.Error(fmt.Sprintf("%s: %v", op, err))
		return nil, err
	}

	if err = foundUser.ComparePasswords(dto.Password); err != nil {
		us.log.Error(fmt.Sprintf("%s: %v", op, err))
		return nil, fmt.Errorf("%s: %w", op, repository.ErrInvalidCredentials)
	}

	// !!! IMPORTANT DO NOT REMOVE !!! This remove password_hash from json struct
	foundUser.SanitizePassword()

	accessToken, _, err := us.jwtManager.GenerateTokens(foundUser)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrInvalidCredentials)
	}

	return &users.UserWithToken{
		User:  foundUser,
		Token: accessToken,
	}, nil
}

func (us *userUsecase) GetByID(ctx context.Context, userID uuid.UUID) (*users.User, error) {
	return us.repo.GetByID(ctx, userID)
}

func (us *userUsecase) GetByUsername(ctx context.Context, username string) (*users.User, error) {
	return us.repo.GetByUsername(ctx, username)
}

func (us *userUsecase) Create(ctx context.Context, user *users.User) error {

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

func (us *userUsecase) Update(ctx context.Context, userID uuid.UUID, dto users.UserUpdateDto) error {
	return us.repo.Update(ctx, userID, dto)
}
