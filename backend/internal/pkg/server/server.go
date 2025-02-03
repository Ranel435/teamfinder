package server

import (
	"encoding/json"
	"net/http"
	"teamfinder/backend/internal/handlers"
	"teamfinder/backend/internal/pkg/storage"
	"teamfinder/backend/internal/routes"
	"teamfinder/backend/internal/services"

	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

type Server struct {
	Router  *gin.Engine
	addr    string
	storage *storage.Storage
}

type Entry struct {
	Value string `json:"value"`
}

func New(host string, st *storage.Storage) *Server {
	s := &Server{
		addr:    host,
		Router:  gin.New(),
		storage: st,
	}

	// CORS configuration
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://teamfinder-hack.ru"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//setup routes
	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	//test routes
	s.Router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World!")
	})
	s.Router.PUT("/key/:key", s.handlerSet)
	s.Router.GET("/key/:key", s.handlerGet)

	//auth routes
	routes.SetupAuthRoutes(s.Router, handlers.NewAuthHandler(services.NewEmailService(), services.NewTelegramService()))

	api := s.Router.Group("/api")

	//hackathon routes
	routes.SetupHackathonRoutes(api, handlers.NewHackathonHandler(services.NewHackathonService()))
	//profile routes
	routes.SetupProfileRoutes(api, handlers.NewProfileHandler(services.NewProfileService()))
}

func (r *Server) handlerSet(ctx *gin.Context) {
	key := ctx.Param("key")

	var v Entry

	if err := json.NewDecoder(ctx.Request.Body).Decode(&v); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	r.storage.Set(key, v.Value)

	ctx.Status(http.StatusOK)
}

func (r *Server) handlerGet(ctx *gin.Context) {
	key := ctx.Param("key")

	v := r.storage.Get(key)
	if v == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, Entry{Value: *v})
}

func (s *Server) Start() {
	s.Router.Run(s.addr)
}

func (s *Server) GetRouter() *gin.Engine {
	return s.Router
}
