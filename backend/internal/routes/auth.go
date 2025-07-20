package routes

import (
	"teamfinder/backend/internal/handlers"
	"teamfinder/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		// Public routes - не требуют аутентификации
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)

		// Email verification routes
		auth.POST("/login/email", authHandler.SendEmailCode)
		auth.POST("/verify/email", authHandler.VerifyEmailCode)

		// Telegram routes
		auth.GET("/login/telegram", authHandler.TelegramLogin)
		auth.POST("/verify/telegram", authHandler.VerifyTelegram)

		// Token routes - refresh не требует аутентификации (использует refresh token)
		auth.POST("/refresh", authHandler.RefreshToken)

		// Protected routes - требуют JWT аутентификации
		protected := auth.Group("")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/check", authHandler.CheckToken)
			protected.POST("/logout", authHandler.Logout)
			protected.DELETE("/account", authHandler.DeleteAccount)
		}
	}
}
