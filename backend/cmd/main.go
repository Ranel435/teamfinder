// TeamFinder API
// Backend service for hackathon team finding platform
//
// @title           TeamFinder API
// @description     Complete backend API for TeamFinder - a platform for finding hackathon teammates
// @version         1.0
// @termsOfService  https://teamfinder-hack.ru/terms
//
// @contact.name   TeamFinder Support
// @contact.email  support@teamfinder-hack.ru
//
// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT
//
// @host      localhost:8090
// @BasePath  /
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
//
// @externalDocs.description  TeamFinder GitHub
// @externalDocs.url          https://github.com/teamfinder/teamfinder
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
