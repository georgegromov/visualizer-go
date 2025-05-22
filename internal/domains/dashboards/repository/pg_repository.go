package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/domains/dashboards"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrVisualizationNotFound                   = errors.New("visualization not found")
	ErrVisualizationsNotFound                  = errors.New("visualizations not found")
	ErrFailedToCreateVisualization             = errors.New("failed to create visualization")
	ErrFailedToUpdateVisualization             = errors.New("failed to update visualization")
	ErrFailedToIncrementViewCountVisualization = errors.New("failed to increment view count visualization")
)

// TODO: УБРАТЬ OP из возврата ошибок

type dashboardRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewDashboardRepo(log *slog.Logger, db *sqlx.DB) dashboards.Repository {
	return &dashboardRepo{log: log, db: db}
}

func (r *dashboardRepo) GetAll(ctx context.Context) ([]*dashboards.Dashboard, error) {
	const op = "repository.dashboardRepo.GetAll"

	dashboards := []*dashboards.Dashboard{}

	query := `
  SELECT 
    d.*, 
    u.username AS creator_name,
    t.name AS template_name
  FROM dashboards d
    LEFT JOIN users u ON d.creator_id = u.id
    LEFT JOIN templates t ON d.template_id = t.id
  ORDER BY d.updated_at DESC;
	`

	err := r.db.SelectContext(ctx, &dashboards, query)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return dashboards, nil
}

func (r *dashboardRepo) GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*dashboards.Dashboard, error) {
	const op = "repository.dashboardRepo.GetByTemplateID"

	dashboards := []*dashboards.Dashboard{}

	query := `SELECT * FROM dashboards WHERE template_id = $1 ORDER BY updated_at DESC;`

	err := r.db.SelectContext(ctx, &dashboards, query, templateID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return dashboards, nil
}

func (r *dashboardRepo) GetByID(ctx context.Context, dashboardId uuid.UUID) (*dashboards.Dashboard, error) {
	const op = "repository.dashboardRepo.GetByID"

	dashboard := &dashboards.Dashboard{}
	err := r.db.GetContext(ctx, dashboard, "SELECT * FROM dashboards WHERE id = $1", dashboardId)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return dashboard, nil
}

func (r *dashboardRepo) GetByShareID(ctx context.Context, shareID uuid.UUID) (*dashboards.Dashboard, error) {
	const op = "repository.dashboardRepo.GetByShareID"

	dashboard := &dashboards.Dashboard{}
	err := r.db.GetContext(ctx, dashboard, "SELECT * FROM dashboards WHERE share_id = $1 AND is_published = TRUE", shareID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return dashboard, nil
}

func (r *dashboardRepo) Create(ctx context.Context, dashboard *dashboards.Dashboard) (uuid.UUID, error) {
	const op = "repository.dashboardRepo.Create"

	var dashboardId uuid.UUID

	err := r.db.GetContext(ctx, &dashboardId, "INSERT INTO dashboards (name, creator_id, template_id) VALUES ($1, $2, $3) RETURNING id",
		dashboard.Name, dashboard.CreatorID, dashboard.TemplateID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return uuid.Nil, err
	}

	return dashboardId, nil
}

func (r *dashboardRepo) Update(ctx context.Context, visualizationID uuid.UUID, dto dashboards.VisualizationUpdateDto) error {
	const op = "repository.dashboardRepo.Update"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	fmt.Println("dto", dto)

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

	if dto.Client != nil {
		setValues = append(setValues, fmt.Sprintf("client=$%d", argId))
		args = append(args, *dto.Client)
		argId++
	}

	if dto.IsPublished != nil {
		setValues = append(setValues, fmt.Sprintf("is_published=$%d", argId))
		args = append(args, *dto.IsPublished)
		argId++
	}

	if dto.TemplateID != nil {
		setValues = append(setValues, fmt.Sprintf("template_id=$%d", argId))
		args = append(args, *dto.TemplateID)
		argId++
	}

	if dto.Tenant != nil {
		setValues = append(setValues, fmt.Sprintf("tenant=$%d", argId))
		args = append(args, *dto.Tenant)
		argId++
	}

	setValues = append(setValues, "updated_at=NOW()")

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE dashboards SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, visualizationID)

	if _, err := r.db.ExecContext(ctx, q, args...); err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return err
	}

	return nil
}

func (r *dashboardRepo) IncrementViewCount(ctx context.Context, visualizationID uuid.UUID) error {
	const op = "repository.dashboardRepo.IncrementViewCount"

	if _, err := r.db.ExecContext(ctx, "UPDATE dashboards SET view_count = view_count + 1, viewed_at=NOW() WHERE id = $1", visualizationID); err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return err
	}

	return nil
}

func (r *dashboardRepo) Delete(ctx context.Context, visualizationID uuid.UUID) error {
	const op = "repository.dashboardRepo.Delete"

	_, err := r.db.ExecContext(ctx, "DELETE FROM dashboards WHERE id = $1", visualizationID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %v", op, err))
		return err
	}

	return nil
}
