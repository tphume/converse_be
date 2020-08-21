package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tphume/converse_be"
)

// Interface for interacting with the users database service
type UserDB interface {
	CreateUser(user *converse_be.User) error
	ReadUserWithCredentials(user *converse_be.User) error
	UpdatePassword(id string, password string) error
	UpdateStatus(id string, status string) error
}

// Executes business logic related to users
type UserService struct {
	DB UserDB
}

// Method to satisfy Handler interface in api.go
func (s *UserService) InitRoute(router *gin.RouterGroup) error {
	group := router.Group("/user")
	group.POST("/signup", s.SignUp)
	group.POST("/login", s.Login)
	group.PUT("/update", s.Update)

	return nil
}

// Below are the methods for our web service api
func (s *UserService) SignUp(ctx *gin.Context) {

}

func (s *UserService) Login(ctx *gin.Context) {

}

// Update is a multiplexer to ChangePassword and ChangeStatus
func (s *UserService) Update(ctx *gin.Context) {

}

func (s *UserService) ChangePassword(ctx *gin.Context) {

}

func (s *UserService) ChangeStatus(ctx *gin.Context) {

}

// Error messages
const (
	errUserNotFound  = "could not find matching user"
	errUserDuplicate = "a user with that name already exist"
)
