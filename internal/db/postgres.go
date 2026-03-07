// folder name become package ---
package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	/*
		conn, err := pgxpool.New(context.Background(),
			"postgres://admin:admin@localhost:5432/apidb")
	*/
	conn, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		return err
	}

	DB = conn
	return nil
}
