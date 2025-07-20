package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB(databaseURL string) error {
	var err error
	maxRetries := 30
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		config, parseErr := pgxpool.ParseConfig(databaseURL)
		if parseErr != nil {
			return fmt.Errorf("failed to parse database URL: %w", parseErr)
		}

		Pool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err == nil {
			if pingErr := Pool.Ping(context.Background()); pingErr == nil {
				log.Println("Successfully connected to database")
				return nil
			} else {
				Pool.Close()
				err = pingErr
			}
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)

		if i < maxRetries-1 {
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	return fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
		log.Println("Database connection pool closed")
	}
}

func GetDB() *pgxpool.Pool {
	return Pool
}
