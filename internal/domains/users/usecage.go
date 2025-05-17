package users

import (
	"context"

	"github.com/google/uuid"
)

type Usecase interface {
	Login(ctx context.Context, dto UserLoginDto) (*UserWithToken, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, userID uuid.UUID, dto UserUpdateDto) error
}
