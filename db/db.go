package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

// Client representing the services in the system that interacts with the db
type Client struct {
	User *UserService
}

// Helper function to return SQL connection
func GetSQLConn(driver string, dataSource string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := sqlx.ConnectContext(ctx, driver, dataSource)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
