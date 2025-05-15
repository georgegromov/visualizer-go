package models

import (
	"time"

	"github.com/google/uuid"
)

type Canvas struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        *string   `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	TemplateID  uuid.UUID `json:"templateId" db:"template_id"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}
