package canvases

import "github.com/google/uuid"

type CanvasCreateDto struct {
	TemplateID uuid.UUID `json:"templateId" db:"template_id" validate:"required"`
}

type CanvasUpdateDto struct {
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}
