package templates

import (
	"time"

	"github.com/google/uuid"
)

type Template struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	IsDeleted   bool       `json:"isDeleted" db:"is_deleted"`
	CreatorID   uuid.UUID  `json:"creatorId" db:"creator_id" validate:"required"`
	Uses        *uint      `json:"uses" db:"uses"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
}
