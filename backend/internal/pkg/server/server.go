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
)

type Server struct {
	Router  *gin.Engine
	addr    string
	storage *storage.Storage
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

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"service":   "teamfinder-backend",
		"version":   "1.0.0",
	})
}

func (s *Server) Start() {
	s.Router.Run(s.addr)
}

func (s *Server) GetRouter() *gin.Engine {
	return s.Router
}
