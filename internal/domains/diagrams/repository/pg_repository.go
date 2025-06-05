package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"visualizer-go/internal/domains/diagrams"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type diagramRepo struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewDiagramRepo(log *slog.Logger, db *sqlx.DB) diagrams.Repository {
	return &diagramRepo{log: log, db: db}
}

func (r *diagramRepo) GetAll(ctx context.Context) ([]*diagrams.Diagram, error) {
	const op = "repository.diagramRepo.GetAll"

	diagrams := []*diagrams.Diagram{}

	q := `select * from diagrams order by updated_at DESC;`

	if err := r.db.SelectContext(ctx, &diagrams, q); err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return diagrams, nil
}

func (r *diagramRepo) GetByID(ctx context.Context, diagramID uuid.UUID) (*diagrams.Diagram, error) {
	const op = "repository.diagramRepo.GetByID"

	diagram := &diagrams.Diagram{}

	q := `select * from diagrams where id = $1;`

	err := r.db.GetContext(ctx, diagram, q, diagramID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return diagram, nil
}
func (r *diagramRepo) Create(ctx context.Context, dto *diagrams.DiagramCreateDTO) (uuid.UUID, error) {
	const op = "repository.diagramRepo.Create"

	var diagramID uuid.UUID

	q := `insert into diagrams (name, content) values ($1, $2) returning id;`

	err := r.db.GetContext(ctx, &diagramID, q, dto.Name, dto.Content)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return uuid.Nil, err
	}

	return diagramID, nil
}

func (r *diagramRepo) Update(ctx context.Context, diagramID uuid.UUID, dto *diagrams.DiagramUpdateDTO) error {
	const op = "repository.diagramRepo.Update"

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if dto.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *dto.Name)
		argId++
	}

	// TODO if content

	setValues = append(setValues, "updated_at=NOW()")

	setQuery := strings.Join(setValues, ", ")

	q := fmt.Sprintf("update diagrams SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, diagramID)

	if _, err := r.db.ExecContext(ctx, q, args...); err != nil {
		r.log.Error(fmt.Sprintf("%s: %s", op, err))
		return err
	}

	return nil
}

func (r *diagramRepo) Delete(ctx context.Context, diagramID uuid.UUID) error {
	const op = "repository.diagramRepo.Delete"

	q := `delete from diagrams where id = $1`

	_, err := r.db.ExecContext(ctx, q, diagramID)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: %v", op, err))
		return err
	}

	return nil
}
