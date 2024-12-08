package main

import (
	"context"
	"log"
	"teamfinder/backend/internal/config"
	"teamfinder/backend/internal/db"
	"teamfinder/backend/internal/pkg/server"
	"teamfinder/backend/internal/pkg/storage"
	"time"

	"github.com/gin-contrib/cors"
)

func main() {
	//==========> DATA BASE (POSTGRESQL)
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db.ConnectDB(cfg.DatabaseURL)
	defer db.Pool.Close()

	// Проверка подключения
	if err := db.Pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	// Запуск миграций
	if err := db.RunMigrations(); err != nil {
		log.Printf("Warning: Failed to run migrations: %v", err)
	}

	//==========> SERVER (ROUTES)

	// store init
	store, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	// router creating
	router := server.New(":8090", &store)

	// CORS configuration
	router.GetRouter().Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:3000"},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Accept", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowWildcard:    true,
	}))

	//start server
	router.Start()
	log.Println("Backend started successfully")
}
