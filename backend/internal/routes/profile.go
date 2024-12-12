package routes

import (
	"teamfinder/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProfileRoutes(router *gin.RouterGroup, handler *handlers.ProfileHandler) {
	profiles := router.Group("/profiles")
	{
		profiles.GET("/hackathon/:id", handler.GetProfilesByHackathonID)
		profiles.POST("", handler.CreateProfile)
		profiles.GET("/:id", handler.GetProfileByID)
		profiles.PUT("/:id", handler.UpdateProfile)
		profiles.DELETE("/:id", handler.DeleteProfile)
	}
}
