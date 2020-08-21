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

func (s *UserService) SignUp(ctx *gin.Context) {

}

func (s *UserService) Login(ctx *gin.Context) {

}

func (s *UserService) ChangePassword(ctx *gin.Context) {

}

func (s *UserService) ChangeStatus(ctx *gin.Context) {

}
