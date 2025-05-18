package dashboards

import (
	"github.com/google/uuid"
)

type DashboardCreateDTO struct {
	TemplateID *uuid.UUID `json:"templateId" db:"template_id"`
}

type VisualizationUpdateDto struct {
	Name        *string    `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	Client      *string    `json:"client" db:"client"`
	IsPublished *bool      `json:"published" db:"is_published"`
	TemplateID  *uuid.UUID `json:"templateId" db:"template_id"`
	Tenant      *string    `json:"tenant" db:"tenant"`
	ViewCount   *uint      `json:"viewCount" db:"view_count"`
}
