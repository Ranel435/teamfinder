package routes

import (
	"teamfinder/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
	// Email routes
	auth.POST("/login/email", authHandler.SendEmailCode)
	auth.POST("/verify/email", authHandler.VerifyEmailCode)

	// Telegram routes
	auth.GET("/login/telegram", authHandler.TelegramLogin)
	auth.POST("/verify/telegram", authHandler.VerifyTelegram)

	// General routes
	auth.POST("/logout", authHandler.Logout)
	auth.POST("/refresh", authHandler.RefreshToken)
	auth.GET("/check", authHandler.CheckToken)
	auth.DELETE("/account", authHandler.DeleteAccount)
	}
}
