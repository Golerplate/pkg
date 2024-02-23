package database_postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewDatabaseConnection(ctx context.Context, cfg *DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
