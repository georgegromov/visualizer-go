package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"visualizer-go/internal/lib/config"
)

const (
	UsersTable         = "users"
	TemplatesTable     = "templates"
	VisualizationTable = "visualizations"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// TODO:
// Separate constructor logic, connect, ping, etc...

func MustConnect(log *slog.Logger, cfg config.Database) *sqlx.DB {
	var driverName = "postgres"
	var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open(driverName, dataSourceName)

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	if err = db.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}

	log.Info("postgres connection successfully established")

	return db
}
