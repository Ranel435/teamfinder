package server

import (
	"encoding/json"
	"net/http"
	"teamfinder/backend/internal/handlers"
	"teamfinder/backend/internal/pkg/storage"
	"teamfinder/backend/internal/routes"
	"teamfinder/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	addr    string
	storage *storage.Storage
}

type Entry struct {
	Value string `json:"value"`
}

func New(host string, st *storage.Storage) *Server {
	s := &Server{
		addr:    host,
		router:  gin.New(),
		storage: st,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	//test routes
	s.router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World!")
	})
	s.router.PUT("/key/:key", s.handlerSet)
	s.router.GET("/key/:key", s.handlerGet)

	//auth routes group
	routes.SetupAuthRoutes(s.router, handlers.NewAuthHandler(services.NewEmailService(), services.NewTelegramService()))
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
	s.router.Run(s.addr)
}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
