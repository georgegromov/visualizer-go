package dto

import "github.com/google/uuid"

type CanvasCreateDto struct {
	TemplateID uuid.UUID `json:"templateId" db:"template_id"`
}

type CanvasUpdateDto struct {
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}
