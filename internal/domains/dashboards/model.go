package dashboards

import (
	"time"

	"github.com/google/uuid"
)

type Dashboard struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	Description  *string    `json:"description" db:"description"`
	IsPublished  bool       `json:"published" db:"is_published"`
	ShareID      uuid.UUID  `json:"shareId" db:"share_id" validate:"required"`
	CreatorID    uuid.UUID  `json:"creatorId" db:"creator_id" validate:"required"`
	TemplateID   *uuid.UUID `json:"templateId" db:"template_id"`
	ViewCount    int        `json:"viewCount" db:"view_count"`
	LastViewedAt *time.Time `json:"lastViewedAt" db:"last_viewed_at"`
	UpdatedAt    *time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
}
