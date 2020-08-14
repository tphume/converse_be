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

// Create a new user
func (s *UserService) CreateUser(user *converse_be.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	_, err := s.DB.NamedExecContext(ctx, userCreate, user)
	if err != nil {
		return err
	}

	return nil
}

// Read a user by id - doesn't return password
func (s *UserService) ReadUser(user *converse_be.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	stmt, err := s.DB.PrepareNamedContext(ctx, userRead)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := stmt.GetContext(ctx, user, user); err != nil {
		return err
	}

	return nil
}

// Update user's password by id
func (s *UserService) UpdatePassword(id string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	_, err := s.DB.ExecContext(ctx, userUpdatePassword, password, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete a user by id
func (s *UserService) DeleteUser(name string, id string) error {
	panic("implement me")
}

// Exec and Query strings
const (
	userCreate         = `INSERT INTO users (id, name, password) VALUES (:id, :name, :password)`
	userRead           = `SELECT (id, name) FROM users WHERE id=:id`
	userUpdatePassword = `UPDATE users SET password=$1 WHERE id=$2`
)
