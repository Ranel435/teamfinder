package main

import (
	"log"
	"os"

	"teamfinder/backend/internal/db"

	"teamfinder/backend/internal/pkg/server"
	"teamfinder/backend/internal/pkg/storage"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	//==========> DB (POSTGRESSQL)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db.ConnectDB(dsn)
	defer db.Pool.Close()

	//==========> SERVER (ROUTES)

	// store init
	store, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	// router creating
<<<<<<< HEAD
	router := server.New(":5174", &store)

	// CORS configuration
	router.GetRouter().Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},
=======
	router := server.New(":8090", &store)

	// CORS configuration
	router.GetRouter().Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8090"},
>>>>>>> cc9b4dfadc52e372e308e866cc8f2be87a37101c
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowWildcard:    true,
	}))

	//start server
	router.Start()
	log.Println("Backend started successfully")
}
