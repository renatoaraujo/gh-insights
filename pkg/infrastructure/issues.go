package infrastructure

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

func (db Database) InsertIssue(ctx context.Context, ID, repositoryID int64, createdAt, closedAt time.Time) {
	conn, err := pgx.ConnectConfig(ctx, db.config)
	if err != nil {
		panic("unable to connect to infrastructure")
	}
	defer conn.Close(ctx)

	sqlStatement := "INSERT INTO issues (id, repository_id, opened_at, closed_at) VALUES ($1, $2, $3, $4)"
	commandTag, err := conn.Exec(ctx, sqlStatement, ID, repositoryID, createdAt, closedAt)
	if err != nil {
		log.Fatal(err)
	}

	if commandTag.RowsAffected() != 1 {
		panic("no rows inserted")
	}
}
