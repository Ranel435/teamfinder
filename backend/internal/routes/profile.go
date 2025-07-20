package routes

import (
	"teamfinder/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProfileRoutes(router *gin.RouterGroup, handler *handlers.ProfileHandler) {
	profiles := router.Group("/profiles")
	{
		profiles.GET("/hackathon/:id", handler.GetProfilesByHackathonID)
		profiles.POST("/hackathon/:id", handler.CreateProfile)
		profiles.GET("/hackathon/:id/profile/:profile_id", handler.GetProfileByID)
		profiles.PUT("/hackathon/:id/profile/:profile_id", handler.UpdateProfile)
		profiles.DELETE("/hackathon/:id/profile/:profile_id", handler.DeleteProfile)
	}
}
