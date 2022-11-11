package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (db Database) GetConnectionPool(ctx context.Context) *pgxpool.Pool {
	err := db.conn.Ping(ctx)
	if err != nil {
		conn, err := pgxpool.NewWithConfig(ctx, db.config)
		if err != nil {
			panic("unable to connect to infrastructure")
		}

		db.conn = conn
	}

	return db.conn
}

type Database struct {
	config *pgxpool.Config
	conn   *pgxpool.Pool
}

func NewDatabase(ctx context.Context, connString string) (*Database, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse infrastructure config: %w", err)
	}

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic("unable to connect to infrastructure")
	}

	return &Database{
		config: config,
		conn:   conn,
	}, nil
}

func (db Database) Close() {
	db.conn.Close()
}
