package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Issue struct {
	ID       int64
	OpenedAt time.Time
	ClosedAt time.Time
}

func (db Database) InsertIssue(ctx context.Context, ID, repositoryID int64, createdAt, closedAt time.Time) {
	sqlStatement := "INSERT INTO issues (id, repository_id, opened_at, closed_at) VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING"
	_, err := db.GetConnectionPool(ctx).Exec(ctx, sqlStatement, ID, repositoryID, createdAt, closedAt)
	if err != nil {
		log.Fatal(err)
	}
}

func (db Database) GetIssues(ctx context.Context) ([]*Issue, error) {
	rows, err := db.GetConnectionPool(ctx).Query(ctx, "select id, opened_at, closed_at from issues")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var issues []*Issue
	for rows.Next() {
		var i *Issue
		err = rows.Scan(&i.ID, &i.OpenedAt, &i.ClosedAt)
		if err != nil {
			return nil, fmt.Errorf("unable to scan. %w", err)
		}
		issues = append(issues, i)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error. %w", rows.Err())
	}

	return issues, nil
}
