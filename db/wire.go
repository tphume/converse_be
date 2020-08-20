//+build wireinject

package db

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// Initialize DB Client
func InitializeClient(database *sqlx.DB) (*Client, error) {
	wire.Build(
		wire.Struct(new(Client), "User"),
		wire.Struct(new(UserService), "DB"),
	)

	return &Client{}, nil
}
