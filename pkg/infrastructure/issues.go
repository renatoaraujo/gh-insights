package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Issue struct {
	ID                int64
	OpenedAt          time.Time
	ClosedAt          time.Time
	TimeOpenedMinutes float64
}

func (db Database) InsertIssue(ctx context.Context, ID, repositoryID int64, title string, number int, state string, createdAt, closedAt time.Time) {
	if closedAt.IsZero() {
		sqlStatement := "INSERT INTO issues (id, repository_id, title, number, state, opened_at) VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (id) DO NOTHING"
		_, err := db.GetConnectionPool(ctx).Exec(ctx, sqlStatement, ID, repositoryID, title, number, state, createdAt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		sqlStatement := "INSERT INTO issues (id, repository_id, title, number, state, opened_at, closed_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO UPDATE SET (state, closed_at) = ($5, $7)"
		_, err := db.GetConnectionPool(ctx).Exec(ctx, sqlStatement, ID, repositoryID, title, number, state, createdAt, closedAt)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (db Database) GetIssuesClosedByMonthAndYear(ctx context.Context, month int, year int) ([]Issue, error) {
	query := `
	SELECT 
		id,
		opened_at,
		closed_at,
		EXTRACT(EPOCH FROM (closed_at -opened_at ))/3600 AS time_opened_minutes
	FROM
	    issues
	WHERE
		EXTRACT(MONTH FROM closed_at) = $1
		AND EXTRACT(YEAR FROM closed_at) = $2
		AND state = 'closed'
	ORDER BY 
	    closed_at
	`

	rows, err := db.GetConnectionPool(ctx).Query(ctx, query, month, year)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var issues []Issue
	for rows.Next() {
		var i Issue
		err = rows.Scan(&i.ID, &i.OpenedAt, &i.ClosedAt, &i.TimeOpenedMinutes)
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
