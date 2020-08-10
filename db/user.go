package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/tphume/converse-be"
	"time"
)

// Represent interactions to the user table in the database
type UserService struct {
	DB *sqlx.DB
}

// Helper constructor method
func NewUserService(driver string, dataSource string) (*UserService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := sqlx.ConnectContext(ctx, driver, dataSource)
	if err != nil {
		return nil, err
	}

	return &UserService{
		DB: conn,
	}, nil
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
