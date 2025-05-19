package measurements

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

type Measurement struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	Content   types.JSONText `json:"content" db:"content"`
	ChartID   uuid.UUID      `json:"chartId" db:"chart_id"`
	UpdatedAt *time.Time     `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
}
