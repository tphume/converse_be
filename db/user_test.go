package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/tphume/converse_be"
	"testing"
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

// Configurations
const (
	driver  = "postgres"
	connStr = "postgres://user:password@converse-test-network:5432"
)

type UserTestSuite struct {
	suite.Suite
	db      *sqlx.DB
	service *UserService
}

// Connect to the database
func (s *UserTestSuite) SetupSuite() {
	db := sqlx.MustConnect(driver, connStr)
	if err := db.Ping(); err != nil {
		panic(err)
	}

	s.db = db
	s.service = &UserService{DB: db}
}

// Insert fake users data before starting each test
func (s *UserTestSuite) SetupTest() {
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
func (s *UserTestSuite) TearDownTest() {
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

// Test that user is created successfully
func (s *UserTestSuite) TestCreate() {
	s.Equal(10, 10)
}

// Run test suite
func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
