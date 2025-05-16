package users

import (
	"context"
	"visualizer-go/internal/dto"

	"github.com/google/uuid"
)

type Repository interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error
}
