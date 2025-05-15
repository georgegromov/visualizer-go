package repository

import (
	"context"
	"fmt"
	"log/slog"
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

func (c *ChartRepo) GetByCanvasID(ctx context.Context, canvasID uuid.UUID) ([]models.Chart, error) {
	const op = "repository.ChartRepo.GetByCanvasID"

	query := `SELECT * FROM charts WHERE canvas_id = $1`

	var charts []models.Chart
	if err := c.db.SelectContext(ctx, &charts, query, canvasID); err != nil {
		c.log.Error(fmt.Sprintf("%s: failed to fetch charts: %v", op, err))
		// TODO: завернуть ошибку в кастомную
		return nil, err
	}

	return charts, nil
}
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
func (c *ChartRepo) Update(ctx context.Context, chartID uuid.UUID, dto dto.ChartUpdateDto) error {
	return nil
}
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
