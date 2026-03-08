// folder name become package ---
package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost" // Default if not set
	}
	port := os.Getenv("DB_PORT")
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		host,
		port,
		os.Getenv("DB_NAME"),
	)
	// fmt.Println("DB CONNECT:", connStr)
	/*
		conn, err := pgxpool.New(context.Background(),
			"postgres://admin:admin@localhost:5432/apidb")
	*/
	// directly connect --
	/*
		conn, err := pgxpool.New(context.Background(), connStr)
		if err != nil {
			return err
		}

		DB = conn
		return nil
	*/

	var err error
	// using retry loop 5 times with 2 sec gap
	for i := 0; i < 5; i++ {
		conn, err := pgxpool.New(context.Background(), connStr)
		if err == nil {
			err = conn.Ping(context.Background())
			// ping - pong returns
			if err == nil {
				// if successfully connected ---
				DB = conn
				return nil
			}
		}
		fmt.Printf("⏳ Waiting for DB... (attempt %d/5 failed: %v)\n", i+1, err)
		time.Sleep(2 * time.Second)
	}
	// All 5 attempts failed
	return fmt.Errorf("failed to connect to database after 5 attempts: %w", err)
}
