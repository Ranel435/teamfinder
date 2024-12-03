package main

import (
	"log"
	"os"

	"teamfinder/backend/internal/db"
	"teamfinder/backend/internal/handlers"
	"teamfinder/backend/internal/middleware"

	"teamfinder/backend/internal/pkg/server"
	"teamfinder/backend/internal/pkg/storage"
	"teamfinder/backend/internal/services"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	//-----POSTGRESSQL-----
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

	//-----ROUTES^SERVER-----

	store, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	router := server.New(":8090", &store)
	ginRouter := router.GetRouter()

	// CORS configuration
	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173/"}, // Change to your frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize services
	emailService := services.NewEmailService()
	telegramService := services.NewTelegramService()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(emailService, telegramService)

	// Public routes
	auth := ginRouter.Group("/auth")
	{
		auth.POST("/login/email", authHandler.SendEmailCode)
		auth.POST("/verify/email", authHandler.VerifyEmailCode)
		auth.GET("/login/telegram", authHandler.TelegramLogin)
		auth.POST("/verify/telegram", authHandler.VerifyTelegram)
		auth.POST("/refresh", authHandler.RefreshToken)
	}

	// Protected routes
	protected := ginRouter.Group("/auth")
	protected.Use(middleware.AuthRequired())
	{
		protected.POST("/logout", authHandler.Logout)
		protected.GET("/check", authHandler.CheckToken)
		protected.DELETE("/account", authHandler.DeleteAccount)
	}

	// router.Run(":8090")
	router.Start()
	log.Println("Application started successfully")
}
