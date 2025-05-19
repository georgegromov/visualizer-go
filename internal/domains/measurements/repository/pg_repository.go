package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/domains/measurements"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type measurementRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewMeasurementRepo(log *slog.Logger, db *sqlx.DB) measurements.Repository {
	return &measurementRepo{log: log, db: db}
}

func (r *measurementRepo) GetByChartID(ctx context.Context, ChartID uuid.UUID) ([]*measurements.Measurement, error) {
	const op = "repository.measurementRepo.GetByChartID"
	m := []*measurements.Measurement{}

	q := `SELECT * FROM measurements WHERE chart_id = $1`

	if err := r.db.SelectContext(ctx, &m, q, ChartID); err != nil {
		r.log.Error(fmt.Sprintf("%s: an error occured while selecting measurements: %v", op, err))
		return nil, err
	}

	return m, nil
}
func (r *measurementRepo) Create(ctx context.Context, dto *measurements.MeasurementCreateDTO) error {
	const op = "repository.measurementRepo.Create"

	q := `INSERT INTO measurements (content, chart_id) VALUES ($1, $2)`

	marshaledContent, err := json.Marshal(dto.Content)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: an error occured while marshalling chart content: %v", op, err))
		return err
	}

	if _, err := r.db.ExecContext(ctx, q, marshaledContent, dto.ChartID); err != nil {
		r.log.Error(fmt.Sprintf("%s: an error occured while inserting charts: %v", op, err))
		return err
	}

	return nil
}
func (r *measurementRepo) Update(ctx context.Context, measurementID uuid.UUID, dto *measurements.MeasurementUpdateDTO) error {
	const op = "repository.measurementRepo.Update"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.Content != nil {
		marshaledMeasurements, err := json.Marshal(dto.Content)
		if err != nil {
			r.log.Error(fmt.Sprintf("%s: an error occured while marshalling chart content: %v", op, err))
			return err
		}
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, marshaledMeasurements)
		argId++
	}

	setValues = append(setValues, "updated_at=NOW()")

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("UPDATE measurements SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, measurementID)

	if _, err := r.db.ExecContext(ctx, q, args...); err != nil {
		r.log.Error(fmt.Sprintf("%s: an error occured while updating measurement: %v", op, err))
		return err
	}

	return nil
}
func (r *measurementRepo) Delete(ctx context.Context, measurementID uuid.UUID) error {
	const op = "repository.measurementRepo.Delete"

	q := `DELETE FROM measurements WHERE id = $1`

	result, err := r.db.ExecContext(ctx, q, measurementID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: an error occured while deleting measurement: %v", op, err))
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.log.Error("failed to get rows affected after delete", slog.String("err", err.Error()))
		return err
	}

	if rowsAffected == 0 {
		r.log.Warn("no canvas found to delete", slog.String("id", measurementID.String()))
	}

	return nil
}
