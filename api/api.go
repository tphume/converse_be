package api

import (
	"github.com/gin-gonic/gin"
)

// Interface for our handler routes
type Handler interface {
	HandlerName() string
	InitRoute(group *gin.RouterGroup) error
}

// Our http web service
type Server struct {
	Handlers []*Handler
}

// Error messages
const (
	errInternal = "something went really wrong here"
)
