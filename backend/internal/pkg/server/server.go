package server

import (
	"net/http"
	"teamfinder/backend/internal/handlers"
	"teamfinder/backend/internal/pkg/storage"
	"teamfinder/backend/internal/routes"
	"teamfinder/backend/internal/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "teamfinder/backend/docs" // Import generated docs
)

type Server struct {
	Router  *gin.Engine
	addr    string
	storage *storage.Storage
}

type HealthResponse struct {
	Service   string `json:"service" example:"teamfinder-backend"`
	Status    string `json:"status" example:"healthy"`
	Timestamp string `json:"timestamp" example:"2024-01-01T12:00:00Z"`
	Version   string `json:"version" example:"1.0.0"`
}

func New(host string, st *storage.Storage) *Server {
	s := &Server{
		addr:    host,
		Router:  gin.New(),
		storage: st,
	}

	s.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://teamfinder-hack.ru"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Swagger route
	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	// Health Check endpoint
	s.Router.GET("/health", s.healthCheck)

	// Auth routes
	routes.SetupAuthRoutes(s.Router, handlers.NewAuthHandler(services.NewEmailService(), services.NewTelegramService()))

	api := s.Router.Group("/api")

	// Hackathon routes
	routes.SetupHackathonRoutes(api, handlers.NewHackathonHandler(services.NewHackathonService()))
	// Profile routes
	routes.SetupProfileRoutes(api, handlers.NewProfileHandler(services.NewProfileService()))
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Check if the service is running and healthy
// @Tags         System
// @Produce      json
// @Success      200 {object} HealthResponse "Service is healthy"
// @Router       /health [get]
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Service:   "teamfinder-backend",
		Status:    "healthy",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "1.0.0",
	})
}

func (s *Server) Start() {
	s.Router.Run(s.addr)
}

func (s *Server) GetRouter() *gin.Engine {
	return s.Router
}
