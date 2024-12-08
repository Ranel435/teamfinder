package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations() error {
	files, err := filepath.Glob("migrations/*.up.sql")
	if err != nil {
		return fmt.Errorf("failed to find migration files: %v", err)
	}

	sort.Strings(files)

	for _, file := range files {
		log.Printf("Running migration: %s", file)

		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", file, err)
		}

		queries := strings.Split(string(content), ";")

		for _, query := range queries {
			query = strings.TrimSpace(query)
			if query == "" {
				continue
			}

			_, err = Pool.Exec(context.Background(), query)
			if err != nil {
				return fmt.Errorf("failed to execute migration %s: %v", file, err)
			}
		}
	}

	return nil
}
