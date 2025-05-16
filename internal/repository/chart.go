package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ChartRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewChartRepo(log *slog.Logger, db *sqlx.DB) *ChartRepo {
	return &ChartRepo{log: log, db: db}
}

// Get By Canvas ID
func (c *ChartRepo) GetByCanvasID(ctx context.Context, canvasID uuid.UUID) ([]*models.Chart, error) {
	const op = "repository.ChartRepo.GetByCanvasID"

	query := `SELECT * FROM charts WHERE canvas_id = $1`

	charts := []*models.Chart{}
	if err := c.db.SelectContext(ctx, &charts, query, canvasID); err != nil {
		c.log.Error(fmt.Sprintf("%s: failed to fetch charts: %v", op, err))
		// TODO: завернуть ошибку в кастомную
		return nil, err
	}

	return charts, nil
}

// Create
func (c *ChartRepo) Create(ctx context.Context, dto dto.ChartCreateDto) error {
	const op = "repository.ChartRepo.Create"

	query := `INSERT INTO charts (type, canvas_id) VALUES ($1, $2)`

	if _, err := c.db.ExecContext(ctx, query, dto.Type, dto.CanvasID); err != nil {
		c.log.Error(fmt.Sprintf("%s: failed to insert chart: %v", op, err))
		// TODO: не возвращать сырую ошибку
		return err
	}

	return nil
}

// Update
func (c *ChartRepo) Update(ctx context.Context, chartID uuid.UUID, dto dto.ChartUpdateDto) error {
	const op = "repository.ChartRepo.Update"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	fmt.Println("dto", dto)

	if dto.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *dto.Name)
		argId++
	}

	if dto.Measurements != nil {
		marshaledMeasurements, err := json.Marshal(dto.Measurements)
		if err != nil {
			c.log.Error(fmt.Sprintf("%s: failed to marshal measurements: %v", op, err))
			return fmt.Errorf("%s: %w", op, ErrFailedToUpdateVisualization)
		}
		setValues = append(setValues, fmt.Sprintf("measurements=$%d", argId))
		args = append(args, marshaledMeasurements)
		argId++
	}

	setValues = append(setValues, "updated_at=NOW()")

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE charts SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, chartID)

	if _, err := c.db.ExecContext(ctx, q, args...); err != nil {
		c.log.Error(fmt.Sprintf("%s: %s", op, err))
		return fmt.Errorf("%w", "failed to update chart")
	}

	return nil
}

// Delete
func (c *ChartRepo) Delete(ctx context.Context, chartID uuid.UUID) error {
	const op = "repository.ChartRepo.Delete"

	query := `DELETE FROM charts WHERE id = $1`

	result, err := c.db.ExecContext(ctx, query, chartID)
	if err != nil {
		c.log.Error(fmt.Sprintf("%s: failed to delete chart: %v", op, err))
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.log.Error("failed to get rows affected after delete", slog.String("err", err.Error()))
		return err
	}

	if rowsAffected == 0 {
		c.log.Warn("no canvas found to delete", slog.String("id", chartID.String()))
	}

	c.log.Info("canvas deleted", slog.String("id", chartID.String()))
	return nil
}
