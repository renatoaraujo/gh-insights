package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Pull struct {
	ID                int64
	State             string
	OpenedAt          time.Time
	ClosedAt          time.Time
	TimeOpenedMinutes float64
}

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

func (db Database) GetOpenedPullsByMonthAndYear(ctx context.Context, month int, year int) ([]Pull, error) {
	query := `
	SELECT 
		id,
		state,
		opened_at
	FROM
	    pulls
	WHERE
		EXTRACT(MONTH FROM opened_at) = $1
		AND EXTRACT(YEAR FROM opened_at) = $2
	ORDER BY 
	    opened_at
	`

	rows, err := db.GetConnectionPool(ctx).Query(ctx, query, month, year)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pulls []Pull
	for rows.Next() {
		var pr Pull
		err = rows.Scan(&pr.ID, &pr.State, &pr.OpenedAt)
		if err != nil {
			return nil, fmt.Errorf("unable to scan. %w", err)
		}
		pulls = append(pulls, pr)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error. %w", rows.Err())
	}

	return pulls, nil
}

func (db Database) GetClosedPullsByMonthAndYear(ctx context.Context, month int, year int) ([]Pull, error) {
	query := `
	SELECT 
		id,
		state,
		closed_at
	FROM
	    pulls
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

	var pulls []Pull
	for rows.Next() {
		var pr Pull
		err = rows.Scan(&pr.ID, &pr.State, &pr.ClosedAt)
		if err != nil {
			return nil, fmt.Errorf("unable to scan. %w", err)
		}
		pulls = append(pulls, pr)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error. %w", rows.Err())
	}

	return pulls, nil
}
