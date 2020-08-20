package db

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/tphume/converse_be"
	"testing"
	"time"
)

// List of fake users to use for testing
var users = []converse_be.User{
	{ID: padID("1"), Username: "Rick", Password: padPassword("RickPassword"), Status: "Ready to get drunk"},
	{ID: padID("2"), Username: "Morty", Password: padPassword("MortyPassword"), Status: "Simping for Jessica"},
	{ID: padID("3"), Username: "Summer", Password: padPassword("SummerPassword"), Status: "Boo-ya"},
	{ID: padID("4"), Username: "Beth", Password: padPassword("BethPassword"), Status: "Going to family therapy"},
	{ID: padID("5"), Username: "Jerry", Password: padPassword("JerryPassword"), Status: "Hi, I'm Jerry"},
}

// Configurations
const (
	driver  = "postgres"
	connStr = "postgres://user:password@psql:5432/?sslmode=disable"
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

// Test Create then read with credentials
func (s *UserTestSuite) TestCreateReadCredentials() {
	user := &converse_be.User{ID: padID("6"), Username: "Meeseeks", Password: padPassword("MeeseeksPassword")}
	s.Require().NoError(s.service.CreateUser(user))

	readUser := &converse_be.User{Username: "Meeseeks"}
	s.Require().NoError(s.service.ReadUserWithCredentials(readUser))

	s.Equal(*user, *readUser)
}

// Test Update password and status
func (s *UserTestSuite) TestUpdate() {
	newPassword := padPassword("newPassword")
	newStatus := "newStatus"

	s.Require().NoError(s.service.UpdatePassword(users[0].ID, newPassword))
	s.Require().NoError(s.service.UpdateStatus(users[0].ID, newStatus))

	readUserCredentials := &converse_be.User{Username: users[0].Username}
	s.Require().NoError(s.service.ReadUserWithCredentials(readUserCredentials))
	s.Equal(newPassword, readUserCredentials.Password)
	s.Equal(newStatus, readUserCredentials.Status)

	readUser := &converse_be.User{ID: padID("1")}
	s.Require().NoError(s.service.ReadUser(readUser))
	s.Equal(newStatus, readUser.Status)
	s.Equal("", readUser.Password)
}

// Test delete and create
func (s *UserTestSuite) TestDeleteCreate() {
	for _, user := range users {
		s.Require().NoError(s.service.DeleteUser(user.ID))
	}

	for _, user := range users {
		s.Require().NoError(s.service.CreateUser(&user))
	}
}

// Run test suite
func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

// Padded generator (since our db uses CHAR type)
func padID(s string) string {
	return fmt.Sprintf("%-36v", s)
}

func padPassword(s string) string {
	return fmt.Sprintf("%-60v", s)
}
