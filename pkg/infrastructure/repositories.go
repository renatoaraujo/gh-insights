package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func (db Database) GetRepos(ctx context.Context) {
	var ID int64
	var Name string
	var URL string

	conn, err := pgx.ConnectConfig(ctx, db.config)
	if err != nil {
		panic("unable to connect to infrastructure")
	}
	defer conn.Close(ctx)

	err = conn.QueryRow(ctx, "select id, name, url from repositories").Scan(&ID, &Name, &URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ID, Name, URL)
}

func (db Database) InsertRepo(ctx context.Context, ID int64, name, url string) {
	conn, err := pgx.ConnectConfig(ctx, db.config)
	if err != nil {
		panic("unable to connect to infrastructure")
	}
	defer conn.Close(ctx)

	sqlStatement := "INSERT INTO repositories (id, name, url) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id"
	_, err = conn.Exec(ctx, sqlStatement, ID, name, url)
	if err != nil {
		log.Fatal(err)
	}

}
