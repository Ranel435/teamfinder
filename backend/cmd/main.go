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
	//---------- DB (POSTGRESSQL)
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

	//---------- SERVER (ROUTES)

	//store init
	store, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	//router creating
	router := server.New(":5173", &store)

	// Updated CORS configuration
	router.GetRouter().Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowWildcard:    true,
	}))

	// // Initialize services
	// emailService := services.NewEmailService()
	// telegramService := services.NewTelegramService()

	// // Initialize handlers
	// authHandler := handlers.NewAuthHandler(emailService, telegramService)

	// // Public routes
	// auth := router.GetRouter().Group("/auth")
	// {
	// 	auth.POST("/login/email", authHandler.SendEmailCode)
	// 	auth.POST("/verify/email", authHandler.VerifyEmailCode)
	// 	auth.GET("/login/telegram", authHandler.TelegramLogin)
	// 	auth.POST("/verify/telegram", authHandler.VerifyTelegram)
	// 	auth.POST("/refresh", authHandler.RefreshToken)
	// }

	// // Protected routes
	// protected := router.GetRouter().Group("/auth")
	// protected.Use(middleware.AuthRequired())
	// {
	// 	protected.POST("/logout", authHandler.Logout)
	// 	protected.GET("/check", authHandler.CheckToken)
	// 	protected.DELETE("/account", authHandler.DeleteAccount)
	// }

	//start server
	// router.Run(":5173")
	router.Start()
	log.Println("Backend started successfully")
}
