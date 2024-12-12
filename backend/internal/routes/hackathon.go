package routes

import (
	"teamfinder/backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupHackathonRoutes(router *gin.RouterGroup, handler *handlers.HackathonHandler) {
	hackathons := router.Group("/hackathons")
	{
		hackathons.GET("", handler.GetAllHackathons)
		hackathons.GET("/:id", handler.GetHackathonByID)
		hackathons.POST("", handler.CreateHackathon)
		hackathons.PUT("/:id", handler.UpdateHackathon)
		hackathons.DELETE("/:id", handler.DeleteHackathon)
	}
}
