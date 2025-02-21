package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strings"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
)

var (
	ErrTemplateNotFound       = errors.New("template not found")
	ErrTemplatesNotFound      = errors.New("templates not found")
	ErrFailedToCreateTemplate = errors.New("failed to create template")
	ErrFailedToUpdateTemplate = errors.New("failed to update template")
)

type TemplateRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewTemplateRepo(log *slog.Logger, db *sqlx.DB) *TemplateRepo {
	return &TemplateRepo{log: log, db: db}
}

func (r *TemplateRepo) GetAll(ctx context.Context) ([]models.Template, error) {
	const op = "repository.TemplateRepo.GetAll"

	var templates []models.Template
	err := r.db.SelectContext(ctx, &templates, "SELECT * FROM templates WHERE is_deleted = FALSE ORDER BY updated_at DESC")
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, ErrTemplatesNotFound)
		}
		return nil, fmt.Errorf("%s: failed to get templates: %w", op, err)
	}

	return templates, nil
}

func (r *TemplateRepo) GetByID(ctx context.Context, templateID uuid.UUID) (models.Template, error) {
	const op = "repository.TemplateRepo.GetByID"

	var template models.Template
	err := r.db.GetContext(ctx, &template, "SELECT * FROM templates WHERE id = $1 AND is_deleted = FALSE", templateID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return template, fmt.Errorf("%s: %w", op, ErrTemplateNotFound)
		}
		return template, fmt.Errorf("%s: failed to get template by ID: %w", op, err)
	}

	return template, nil
}

func (r *TemplateRepo) Create(ctx context.Context, dto dto.TemplateCreateDto) error {
	const op = "repository.TemplateRepo.Create"

	_, err := r.db.ExecContext(ctx, "INSERT INTO templates (name, description) VALUES ($1, $2)",
		dto.Name, dto.Description)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return fmt.Errorf("%s: %w", op, ErrFailedToCreateTemplate)
	}

	return nil
}

func (r *TemplateRepo) Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error {
	const op = "repository.TemplateRepo.Update"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *dto.Name)
		argId++
	}

	if dto.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *dto.Description)
		argId++
	}

	if dto.Canvases != nil {
		setValues = append(setValues, fmt.Sprintf("canvases=$%d", argId))
		args = append(args, *dto.Canvases)
		argId++
	}

	if dto.IsDeleted != nil {
		setValues = append(setValues, fmt.Sprintf("is_deleted=$%d", argId))
		args = append(args, *dto.IsDeleted)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE templates SET %s WHERE id=$%d AND is_deleted = FALSE", setQuery, argId)
	args = append(args, templateID)

	if _, err := r.db.ExecContext(ctx, q, args...); err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return fmt.Errorf("%s: %w", op, ErrFailedToUpdateTemplate)
	}

	return nil
}
