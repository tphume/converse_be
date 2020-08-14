package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/tphume/converse_be"
	"time"
)

// Represent interactions to the user table in the database
type UserService struct {
	DB *sqlx.DB
}

// Basic CRUD methods
func (s *UserService) CreateUser(user *converse_be.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	_, err := s.DB.NamedExecContext(ctx, sqlCreate, user)
	if err != nil {
		return err
	}

	return nil
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

// Exec and Query strings
const (
	sqlCreate = `INSERT INTO users (id, name, password) VALUES (:id, :name, :password)`
)
