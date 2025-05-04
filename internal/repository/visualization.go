package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"

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

type VisualizationRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewVisualizationRepo(log *slog.Logger, db *sqlx.DB) *VisualizationRepo {
	return &VisualizationRepo{log: log, db: db}
}

func (r *VisualizationRepo) GetAll(ctx context.Context) ([]models.Visualization, error) {
	const op = "repository.VisualizationRepo.GetAll"

	var visualizations []models.Visualization

	query := `
	SELECT 
			v.id, 
			v.name, 
			v.description,
			v.client,
			v.is_published, 
			v.share_id,
      v.template_id,
			v.updated_at, 
			v.created_at, 
      v.view_count,
      v.viewed_at,
			v.user_id,
			u.username AS username,
      t.name AS template_name
	FROM visualizations v
	LEFT JOIN users u ON v.user_id = u.id
  LEFT JOIN templates t ON v.template_id = t.id
	ORDER BY v.updated_at DESC
	`

	err := r.db.SelectContext(ctx, &visualizations, query)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w", ErrVisualizationsNotFound)
		}
		return nil, fmt.Errorf("failed to get visualizations")
	}

	return visualizations, nil
}

func (r *VisualizationRepo) GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]models.Visualization, error) {
	const op = "repository.VisualizationRepo.GetByTemplateID"

	var visualizations []models.Visualization

	query := `
  SELECT 
    id, 
    name
  FROM visualizations
  WHERE template_id = $1
  ORDER BY updated_at DESC;
  `

	err := r.db.SelectContext(ctx, &visualizations, query, templateID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w", ErrVisualizationsNotFound)
		}
		return nil, fmt.Errorf("failed to get visualizations")
	}

	return visualizations, nil
}

func (r *VisualizationRepo) GetByID(ctx context.Context, visualizationID uuid.UUID) (models.Visualization, error) {
	const op = "repository.VisualizationRepo.GetByID"

	var visualization models.Visualization
	err := r.db.GetContext(ctx, &visualization, "SELECT * FROM visualizations WHERE id = $1", visualizationID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return visualization, fmt.Errorf("%w", ErrVisualizationNotFound)
		}
		return visualization, fmt.Errorf("failed to get visualization by ID")
	}

	return visualization, nil
}

func (r *VisualizationRepo) GetByShareID(ctx context.Context, shareID uuid.UUID) (models.Visualization, error) {
	const op = "repository.VisualizationRepo.GetByShareID"

	var visualization models.Visualization
	err := r.db.GetContext(ctx, &visualization, "SELECT * FROM visualizations WHERE share_id = $1 AND is_published = TRUE", shareID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return visualization, fmt.Errorf("%w", ErrVisualizationNotFound)
		}
		return visualization, fmt.Errorf("failed to get visualization")
	}
	return visualization, nil
}

func (r *VisualizationRepo) Create(ctx context.Context, dto dto.VisualizationCreateDto) (uuid.UUID, error) {
	const op = "repository.VisualizationRepo.Create"

	var visualizationID uuid.UUID

	// Преобразование поля Canvases в JSON
	var canvasesJson interface{}
	var err error

	// Если Canvases не nil, сериализуем в JSON
	if dto.Canvases != nil {
		canvasesJson, err = json.Marshal(dto.Canvases)
		if err != nil {
			r.log.Error(fmt.Sprintf("%s: failed to marshal canvases: %v", op, err))
			return uuid.Nil, fmt.Errorf("%s: %w", op, ErrFailedToCreateVisualization)
		}
	} else {
		// Если Canvases равно nil, передаем NULL
		canvasesJson = nil
	}

	// Вставка данных в таблицу visualizations
	err = r.db.GetContext(ctx, &visualizationID, "INSERT INTO visualizations (name, user_id, canvases, template_id) VALUES ($1, $2, $3, $4) RETURNING id",
		dto.Name, dto.UserID, canvasesJson, dto.TemplateID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return uuid.Nil, fmt.Errorf("%s: %w", op, ErrFailedToCreateVisualization)
	}

	return visualizationID, nil
}

func (r *VisualizationRepo) Update(ctx context.Context, visualizationID uuid.UUID, dto dto.VisualizationUpdateDto) error {
	const op = "repository.VisualizationRepo.Update"

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

	if dto.Canvases != nil {
		canvasesJson, err := json.Marshal(dto.Canvases)
		if err != nil {
			r.log.Error(fmt.Sprintf("%s: failed to marshal canvases: %v", op, err))
			return fmt.Errorf("%s: %w", op, ErrFailedToUpdateVisualization)
		}
		setValues = append(setValues, fmt.Sprintf("canvases=$%d", argId))
		args = append(args, canvasesJson)
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

	// if dto.ViewCount != nil {
	// 	setValues = append(setValues, fmt.Sprintf("view_count=$%d", argId))
	// 	args = append(args, *dto.ViewCount)
	// 	argId++
	// }

	setValues = append(setValues, fmt.Sprintf("is_saved=$%d", argId))
	args = append(args, true)
	argId++

	setValues = append(setValues, fmt.Sprintf("is_publishable=$%d", argId))
	args = append(args, true)
	argId++

	setValues = append(setValues, "updated_at=NOW()")

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE visualizations SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, visualizationID)

	if _, err := r.db.ExecContext(ctx, q, args...); err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return fmt.Errorf("%w", ErrFailedToUpdateVisualization)
	}

	return nil
}

func (r *VisualizationRepo) IncrementViewCount(ctx context.Context, visualizationID uuid.UUID) error {
	const op = "repository.VisualizationRepo.IncrementViewCount"

	if _, err := r.db.ExecContext(ctx, "UPDATE visualizations SET view_count = view_count + 1, viewed_at=NOW() WHERE id = $1", visualizationID); err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return fmt.Errorf("%w", ErrFailedToIncrementViewCountVisualization)
	}

	return nil
}

func (r *VisualizationRepo) Delete(ctx context.Context, visualizationID uuid.UUID) error {
	const op = "repository.VisualizationRepo.Delete"

	_, err := r.db.ExecContext(ctx, "DELETE FROM visualizations WHERE id = $1", visualizationID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %v", op, err))
		return fmt.Errorf("failed to delete visualization")
	}

	return nil
}
