package diagrams

import "github.com/jmoiron/sqlx/types"

type DiagramCreateDTO struct {
	Name    string         `json:"name" db:"name" validate:"required"`
	Content types.JSONText `json:"content" db:"content" validate:"required"`
}

type DiagramUpdateDTO struct {
	Name    *string         `json:"name" db:"name"`
	Content *types.JSONText `json:"content" db:"content"`
}
