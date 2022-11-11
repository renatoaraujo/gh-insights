package infrastructure

import (
	"context"
	"fmt"
	"log"
)

func (db Database) GetRepos(ctx context.Context) {
	var ID int64
	var Name string
	var URL string

	err := db.GetConnectionPool(ctx).QueryRow(ctx, "select id, name, url from repositories").Scan(&ID, &Name, &URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ID, Name, URL)
}

func (db Database) InsertRepo(ctx context.Context, ID int64, name, url string) {
	sqlStatement := "INSERT INTO repositories (id, name, url) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING"
	_, err := db.GetConnectionPool(ctx).Exec(ctx, sqlStatement, ID, name, url)
	if err != nil {
		log.Fatal(err)
	}
}
