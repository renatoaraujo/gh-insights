package infrastructure

import (
	"context"
	"log"
	"time"
)

func (db Database) InsertPull(ctx context.Context, ID, repositoryID int64, title string, number int, state string, createdAt, closedAt time.Time) {
	if closedAt.IsZero() {
		sqlStatement := "INSERT INTO pulls (id, repository_id, title, number, state, opened_at) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (id) DO NOTHING"
		_, err := db.GetConnectionPool(ctx).Exec(ctx, sqlStatement, ID, repositoryID, title, number, state, createdAt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		sqlStatement := "INSERT INTO pulls (id, repository_id, title, number, state, opened_at, closed_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO UPDATE SET closed_at = $7"
		_, err := db.GetConnectionPool(ctx).Exec(ctx, sqlStatement, ID, repositoryID, title, number, state, createdAt, closedAt)
		if err != nil {
			log.Fatal(err)
		}
	}
}
