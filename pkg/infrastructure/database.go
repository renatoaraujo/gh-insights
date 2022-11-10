package infrastructure

import (
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Connection struct {
	Conn *pgx.Conn
}

type Database struct {
	config *pgx.ConnConfig
}

func NewDatabase(connString string) (*Database, error) {
	config, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse infrastructure config: %w", err)
	}

	return &Database{
		config: config,
	}, nil
}
