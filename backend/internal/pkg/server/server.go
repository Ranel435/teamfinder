package server

import (
	"encoding/json"
	"net/http"
	"teamfinder/backend/internal/pkg/storage"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	addr    string
	storage *storage.Storage
}

type Entry struct {
	Value string
}

func New(host string, st *storage.Storage) *Server {
	s := &Server{
		addr:   host,
		router: gin.New(),
	}

	return s
}

func (r *Server) newAPI() *gin.Engine {
	engine := gin.New()

	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World! ")
	})

	engine.PUT("/key/:key", r.handlerSet)
	engine.GET("/key/:key", r.handlerGet)

	return engine
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

func (r *Server) Start() {
	r.newAPI().Run(r.addr)
}

func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
