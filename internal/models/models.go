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
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

type Template struct {
	ID          uuid.UUID       `json:"id" db:"id"`
	Name        string          `json:"name" db:"name"`
	Description *string         `json:"description" db:"description"`
	Canvases    *types.JSONText `json:"canvases" db:"canvases"`
	IsDeleted   bool            `json:"isDeleted" db:"is_deleted"`
	Uses        *uint           `json:"uses" db:"uses"`
	UpdatedAt   time.Time       `json:"updatedAt" db:"updated_at"`
	CreatedAt   time.Time       `json:"createdAt" db:"created_at"`
}

type Visualization struct {
	ID            uuid.UUID       `json:"id" db:"id"`
	Name          string          `json:"name" db:"name"`
	Description   *string         `json:"description" db:"description"`
	Client        *string         `json:"client" db:"client"`
	IsPublished   bool            `json:"published" db:"is_published"`
	ShareID       uuid.UUID       `json:"shareId" db:"share_id"`
	UpdatedAt     time.Time       `json:"updatedAt" db:"updated_at"`
	CreatedAt     time.Time       `json:"createdAt" db:"created_at"`
	UserID        uuid.UUID       `json:"userId" db:"user_id"`
	TemplateID    *uuid.UUID      `json:"templateId" db:"template_id"`
	TemplateName  *string         `json:"templateName" db:"template_name"`
	Canvases      *types.JSONText `json:"canvases" db:"canvases"`
	IsSaved       bool            `json:"saved" db:"is_saved"`
	IsPublishable bool            `json:"publishable" db:"is_publishable"`
	Tenant        *string         `json:"tenant" db:"tenant"`
	Username      *string         `json:"username" db:"username"`
	ViewCount     int             `json:"viewCount" db:"view_count"`
	ViewedAt      *time.Time      `json:"viewedAt" db:"viewed_at"`
}
