package diagrams

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

type Diagram struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Content   types.JSONText `json:"content" db:"content"`
	UpdatedAt *time.Time     `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
}
