package api

import (
	"github.com/gin-gonic/gin"
)

// Interface for our handler routes
type Handler interface {
	InitRoute(router *gin.RouterGroup) error
}

// Our http web service
type Server struct {
	Engine   *gin.Engine
	Handlers []Handler
}

func InitServer(s *Server) {
	router := s.Engine.Group("/api")
	for _, h := range s.Handlers {
		if err := h.InitRoute(router); err != nil {
			panic(err)
		}
	}
}

// Error messages
const (
	errInternal = "something went really wrong here"
)
