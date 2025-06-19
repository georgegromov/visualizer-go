package dashboards

import (
	"time"

	"github.com/google/uuid"
)

type Dashboard struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	Name         string     `json:"name" db:"name"`
	Description  *string    `json:"description" db:"description"`
	Tenant       string     `json:"tenant" db:"tenant"`
	IsPublished  bool       `json:"published" db:"is_published"`
	ShareID      uuid.UUID  `json:"shareId" db:"share_id"`
	CreatorID    uuid.UUID  `json:"creatorId" db:"creator_id"`
	CreatorName  *string    `json:"creatorName" db:"creator_name"`
	TemplateID   *uuid.UUID `json:"templateId" db:"template_id"`
	TemplateName *string    `json:"templateName" db:"template_name"`
	ViewCount    int        `json:"viewCount" db:"view_count"`
	LastViewedAt *time.Time `json:"lastViewedAt" db:"last_viewed_at"`
	UpdatedAt    *time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
}

// type Dashboard struct {
// 	ID           uuid.UUID       `json:"id" db:"id"`
// 	Name         string          `json:"name" db:"name"`
// 	Description  *string         `json:"description" db:"description"`
// 	IsPublished  bool            `json:"published" db:"is_published"`
// 	ShareID      uuid.UUID       `json:"shareId" db:"share_id"`
// 	Creator      CreatorDetails  `json:"creator"`
// 	Template     TemplateDetails `json:"template"`
// 	ViewCount    int             `json:"viewCount" db:"view_count"`
// 	LastViewedAt *time.Time      `json:"lastViewedAt" db:"last_viewed_at"`
// 	UpdatedAt    *time.Time      `json:"updatedAt" db:"updated_at"`
// 	CreatedAt    time.Time       `json:"createdAt" db:"created_at"`
// }

// type CreatorDetails struct {
// 	ID   uuid.UUID `json:"id"`
// 	Name *string    `json:"name"`
// }

// type TemplateDetails struct {
// 	ID       *uuid.UUID      `json:"id" db:"template_id"`
// 	Name     *string         `json:"name" db:"template_name"`
// 	Canvases []CanvasDetails `json:"canvases"`
// }

// type CanvasDetails struct {
// 	ID          uuid.UUID      `json:"caid" db:"canvas_id"`
// 	Name        *string        `json:"name" db:"canvas_name"`
// 	Description *string        `json:"description" db:"canvas_description"`
// 	Charts      []ChartDetails `json:"charts"`
// }
// type ChartDetails struct {
// 	ID   uuid.UUID `json:"id" db:"chart_id"`
// 	Name *string   `json:"name" db:"chart_name"`
// 	Type string    `json:"type" db:"chart_type"`
// }
