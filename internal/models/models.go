package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type Template struct {
	ID          uuid.UUID       `json:"id" db:"id"`
	Name        string          `json:"name" db:"name"`
	Description *string         `json:"description" db:"description"`
	Canvases    *types.JSONText `json:"canvases" db:"canvases"`
	IsDeleted   bool            `json:"is_deleted" db:"is_deleted"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
}

type Visualization struct {
	ID          uuid.UUID       `json:"id" db:"id"`
	Name        string          `json:"name" db:"name"`
	Description *string         `json:"description" db:"description"`
	Client      *string         `json:"client" db:"client"`
	IsPublished bool            `json:"is_published" db:"is_published"`
	ShareID    uuid.UUID       `json:"share_id" db:"share_id"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
	UserID      uuid.UUID       `json:"user_id" db:"user_id"`
	TemplateID  *uuid.UUID      `json:"template_id" db:"template_id"`
	Canvases    *types.JSONText `json:"canvases" db:"canvases"`
	IsSaved bool `json:"is_saved" db:"is_saved"`
	IsPublishable bool `json:"is_publishable" db:"is_publishable"`
	Tenant *string `json:"tenant" db:"tenant"`
	Username    *string         `json:"username" db:"username"`
}