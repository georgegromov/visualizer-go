package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/domains/templates"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrTemplateNotFound       = errors.New("template not found")
	ErrTemplatesNotFound      = errors.New("templates not found")
	ErrFailedToCreateTemplate = errors.New("failed to create template")
	ErrFailedToUpdateTemplate = errors.New("failed to update template")
)

type templateRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewTemplateRepo(log *slog.Logger, db *sqlx.DB) templates.Repository {
	return &templateRepo{log: log, db: db}
}

func (r *templateRepo) GetAll(ctx context.Context) ([]*templates.Template, error) {
	const op = "repository.TemplateRepo.GetAll"

	templates := []*templates.Template{}

	q := `
  SELECT 
    t.id,
    t.name,
    t.description,
    t.is_deleted,
    t.updated_at,
    t.created_at,
    COUNT(DISTINCT d.id) AS uses
  FROM 
    templates t
  LEFT JOIN 
    dashboards d ON d.template_id = t.id
  WHERE 
    t.is_deleted = false
  GROUP BY 
    t.id, t.name, t.description, t.is_deleted, t.updated_at, t.created_at
  ORDER BY 
    t.updated_at DESC;
  `

	err := r.db.SelectContext(ctx, &templates, q)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return templates, nil
}

func (r *templateRepo) GetByID(ctx context.Context, templateID uuid.UUID) (*templates.Template, error) {
	const op = "repository.TemplateRepo.GetByID"

	template := &templates.Template{}
	err := r.db.GetContext(ctx, template, "SELECT * FROM templates WHERE id = $1 AND is_deleted = FALSE", templateID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return template, nil
}

func (r *templateRepo) Create(ctx context.Context, template *templates.Template) (uuid.UUID, error) {
	const op = "repository.TemplateRepo.Create"

	var templateID uuid.UUID

	err := r.db.GetContext(ctx, &templateID, "INSERT INTO templates (name, description, creator_id) VALUES ($1, $2, $3) RETURNING id",
		template.Name, template.Description, template.CreatorID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return uuid.Nil, err
	}

	return templateID, nil
}

func (r *templateRepo) Update(ctx context.Context, templateID uuid.UUID, dto *templates.TemplateUpdateDto) error {
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

	if dto.IsDeleted != nil {
		setValues = append(setValues, fmt.Sprintf("is_deleted=$%d", argId))
		args = append(args, *dto.IsDeleted)
		argId++
	}

	setValues = append(setValues, "updated_at=NOW()")

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE templates SET %s WHERE id=$%d AND is_deleted = FALSE", setQuery, argId)
	args = append(args, templateID)

	if _, err := r.db.ExecContext(ctx, q, args...); err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return err
	}

	return nil
}
