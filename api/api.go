package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Interface for our handler routes
type Handler interface {
	InitRoute(router *gin.RouterGroup) error
}

// Our http web service
type Server struct {
	Address  string
	Engine   *gin.Engine
	Handlers []Handler
}

func (s *Server) InitServer() {
	router := s.Engine.Group("/api")
	for _, h := range s.Handlers {
		if err := h.InitRoute(router); err != nil {
			panic(err)
		}
	}
}

func (s *Server) RunServer() {
	log.Fatal(s.Engine.Run(s.Address))
}

// Error messages
const (
	errInternal = "something went really wrong here"
)
