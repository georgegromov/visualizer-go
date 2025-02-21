package repository

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type (
	Repository struct{}
)

func New(log *slog.Logger, db *sqlx.DB) *Repository {
	return &Repository{}
}
