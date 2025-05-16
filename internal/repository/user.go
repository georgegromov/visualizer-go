package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrFailedToCreateUser = errors.New("failed to create user")
	ErrFailedToUpdateUser = errors.New("failed to update user")
	ErrUserNotFound       = errors.New("user not found")
	ErrFailedToFetchUsers = errors.New("failed to fetch users")
	ErrFailedToLogin      = errors.New("failed to login")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewUserRepo(log *slog.Logger, db *sqlx.DB) *UserRepo {
	return &UserRepo{log: log, db: db}
}

func (r *UserRepo) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	const op = "repository.UserRepo.GetByID"

	user := &models.User{}
	err := r.db.GetContext(ctx, user, "SELECT id, username, role, created_at, updated_at FROM users WHERE id=$1", userID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %v", op, err))
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, ErrFailedToFetchUsers)
	}

	return user, nil
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	const op = "repository.UserRepo.GetByUsername"

	user := &models.User{}
	err := r.db.GetContext(ctx, user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %v", op, err))
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("%w", ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, ErrFailedToFetchUsers)
	}

	return user, nil
}

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
	const op = "repository.UserRepo.Create"

	_, err := r.db.ExecContext(ctx, "INSERT INTO users (username, password_hash) VALUES ($1, $2)",
		user.Username, user.PasswordHash)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %w", op, ErrFailedToCreateUser)
	}

	return nil
}

func (r *UserRepo) Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error {
	const op = "repository.UserRepo.Update"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.Role != nil {
		setValues = append(setValues, fmt.Sprintf("role=$%d", argId))
		args = append(args, *dto.Role)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, userID)

	if _, err := r.db.ExecContext(ctx, q, args...); err != nil {
		r.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("%s: %w", op, ErrFailedToUpdateUser)
	}

	return nil
}
