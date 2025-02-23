package dto

import (
	"github.com/google/uuid"
)

type VisualizationCreateDto struct {
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Client      *string `json:"client" db:"client"`

	UserID uuid.UUID `json:"user_id" db:"user_id"`
}

type VisualizationUpdateDto struct {
	Name        *string      `json:"name" db:"name"`
	Description *string      `json:"description" db:"description"`
	Client      *string      `json:"client" db:"client"`
	IsPublished *bool        `json:"is_published" db:"is_published"`
	Canvases    *interface{} `json:"canvases" db:"canvases"`
}
