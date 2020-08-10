package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/tphume/converse-be"
)

// Represent interactions to the user table in the database
type UserService struct {
	DB *sqlx.DB
}

// Basic CRUD methods
func (s *UserService) CreateUser(user *converse_be.User) error {
	panic("implement me")
}

func (s *UserService) ReadUser(user *converse_be.User) error {
	panic("implement me")
}

func (s *UserService) UpdateUser(user *converse_be.User) error {
	panic("implement me")
}

func (s *UserService) DeleteUser(name string, id string) error {
	panic("implement me")
}
