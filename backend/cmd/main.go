package main

import (
	"teamfinder/backend/internal/handlers"
	"teamfinder/backend/internal/middleware"

	"teamfinder/backend/internal/pkg/server"
	"teamfinder/backend/internal/pkg/storage"
	"teamfinder/backend/internal/services"
)

func main() {
	store, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}

	router := server.New(":8090", &store)
	ginRouter := router.GetRouter()

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
}
