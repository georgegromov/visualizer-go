package dashboards

import (
	"github.com/google/uuid"
)

type DashboardCreateDTO struct {
	TemplateID *uuid.UUID `json:"templateId" db:"template_id"`
}

type DashboardUpdateDto struct {
	Name        *string    `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	IsPublished *bool      `json:"published" db:"is_published"`
	TemplateID  *uuid.UUID `json:"templateId" db:"template_id"`
	Tenant      *string    `json:"tenant" db:"tenant"`
	ViewCount   *uint      `json:"viewCount" db:"view_count"`
}
