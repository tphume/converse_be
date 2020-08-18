package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"github.com/tphume/converse_be"
	"time"
)

// List of fake users to use for testing
var users = []converse_be.User{
	{ID: "1", Username: "Rick", Password: "RickPassword", Status: "Ready to get drunk"},
	{ID: "2", Username: "Morty", Password: "MortyPassword", Status: "Simping for Jessica"},
	{ID: "3", Username: "Summer", Password: "SummerPassword", Status: "Boo-ya"},
	{ID: "4", Username: "Beth", Password: "BethPassword", Status: "Going to family therapy"},
	{ID: "5", Username: "Jerry", Password: "JerryPassword", Status: "Hi, I'm Jerry"},
}

type UserTestSuite struct {
	suite.Suite
	db      *sqlx.DB
	service *UserService
}

// Insert fake users data before starting each test
func (s *UserTestSuite) SetupTestSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	s.Require().NoError(err)

	for _, user := range users {
		_, err = tx.Exec(
			`INSERT INTO users (id, username, password, status) VALUES ($1, $2, $3, $4)`,
			user.ID, user.Username, user.Password, user.Status,
		)

		s.Require().NoError(err)
	}

	s.Require().NoError(tx.Commit())
}

// Delete all fake users before starting each test
func (s *UserTestSuite) TearDownTestSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	s.Require().NoError(err)

	for _, user := range users {
		_, err = tx.Exec(
			`DELETE FROM users WHERE id=$1`,
			user.ID,
		)

		s.Require().NoError(err)
	}

	s.Require().NoError(tx.Commit())
}
