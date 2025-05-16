package repository

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/domains/canvases"
	"visualizer-go/internal/dto"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type canvasRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewCanvasRepo(log *slog.Logger, db *sqlx.DB) canvases.Repository {
	return &canvasRepo{log: log, db: db}
}

func (c *canvasRepo) Create(ctx context.Context, dto dto.CanvasCreateDto) error {
	const op = "repository.CanvasRepo.Create"

	query := `INSERT INTO canvases (template_id) VALUES ($1)`

	if _, err := c.db.ExecContext(ctx, query, dto.TemplateID); err != nil {
		c.log.Error(fmt.Sprintf("%s: failed to insert canvas: %v", op, err))
		// TODO: не возвращать сырую ошибку
		return err
	}

	return nil
}

func (c *canvasRepo) GetCanvasesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*canvases.Canvas, error) {
	const op = "repository.CanvasRepo.GetCanvasesByTemplateID"

	query := `SELECT * FROM canvases WHERE template_id = $1`

	canvases := []*canvases.Canvas{}
	if err := c.db.SelectContext(ctx, &canvases, query, templateID); err != nil {
		c.log.Error(fmt.Sprintf("%s: failed to fetch canvases: %v", op, err))
		// TODO: завернуть ошибку в кастомную
		return nil, err
	}

	return canvases, nil
}

func (c *canvasRepo) Update(ctx context.Context, canvasID uuid.UUID, dto dto.CanvasUpdateDto) error {
	const op = "repository.CanvasRepo.Update"

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

	if len(setValues) == 0 {
		c.log.Warn(fmt.Sprintf("%s: nothing to update", op), slog.String("id", canvasID.String()))
		return nil
	}

	setValues = append(setValues, "updated_at=NOW()")
	setClause := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE canvases SET %s WHERE id=$%d", setClause, argId)
	args = append(args, canvasID)

	if _, err := c.db.ExecContext(ctx, query, args...); err != nil {
		c.log.Error(fmt.Sprintf("%s: failed to update canvas: %v", op, err))
		return fmt.Errorf("%s: %w", op, err)
	}

	c.log.Info("canvas updated", slog.String("id", canvasID.String()))
	return nil
}

func (c *canvasRepo) Delete(ctx context.Context, canvasID uuid.UUID) error {
	const op = "repository.CanvasRepo.Delete"

	query := `DELETE FROM canvases WHERE id = $1`

	result, err := c.db.ExecContext(ctx, query, canvasID)
	if err != nil {
		c.log.Error(fmt.Sprintf("%s: failed to delete canvas: %v", op, err))
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.log.Error("failed to get rows affected after delete", slog.String("err", err.Error()))
		return err
	}

	if rowsAffected == 0 {
		c.log.Warn("no canvas found to delete", slog.String("id", canvasID.String()))
	}

	c.log.Info("canvas deleted", slog.String("id", canvasID.String()))
	return nil
}
