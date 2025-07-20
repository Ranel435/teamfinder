package main

import (
	"context"
	"log"
	"teamfinder/backend/internal/config"
	"teamfinder/backend/internal/database"
	"teamfinder/backend/internal/middleware"
	"teamfinder/backend/internal/pkg/server"
	"teamfinder/backend/internal/pkg/storage"
)

func main() {
	cfg := config.LoadConfig()

	if err := database.ConnectDB(cfg.DatabaseURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.CloseDB()

	if err := database.Pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	if err := database.RunMigrations(); err != nil {
		log.Printf("Warning: Failed to run migrations: %v", err)
	}

	store, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	router := server.New(":8090", &store)
	router.Router.Use(middleware.CorsMiddleware())

	router.Start()
	log.Println("Backend started successfully")
}
