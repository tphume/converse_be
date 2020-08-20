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

// Read a user by username - return password
func (s *UserService) ReadUserWithCredentials(user *converse_be.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	stmt, err := s.DB.PrepareNamedContext(ctx, userCredentials)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := stmt.GetContext(ctx, user, user); err != nil {
		return err
	}

	return nil
}

// Read a user by id - doesn't return password
func (s *UserService) ReadUser(user *converse_be.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := s.DB.QueryRowContext(ctx, userRead, user.ID).Scan(&user.Username, &user.Status); err != nil {
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

// Update a user's status by id
func (s *UserService) UpdateStatus(id string, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	_, err := s.DB.ExecContext(ctx, userUpdateStatus, status, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete a user by id
func (s *UserService) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()

	_, err := s.DB.ExecContext(ctx, userDelete, id)
	if err != nil {
		return err
	}

	return nil
}

// Exec and Query strings
const (
	userCreate         = `INSERT INTO users (id, username, password) VALUES (:id, :username, :password)`
	userCredentials    = `SELECT * FROM users WHERE username=:username LIMIT 1`
	userRead           = `SELECT username, status FROM users WHERE id=$1 LIMIT 1`
	userUpdatePassword = `UPDATE users SET password=$1 WHERE id=$2`
	userUpdateStatus   = `UPDATE users SET status=$1 WHERE id=$2`
	userDelete         = `DELETE FROM users WHERE id=$1`
)
