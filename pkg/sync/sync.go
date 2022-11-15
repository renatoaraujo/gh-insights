package sync

import (
	"context"
	"log"

	"renatoaraujo/gh-insights/pkg/github"
	"renatoaraujo/gh-insights/pkg/infrastructure"
)

func Sync(ctx context.Context, client github.GitHub, db *infrastructure.Database) {
	log.Printf("Starting to sync %s/%s", client.Owner, client.Repo)

	repo, err := client.GetRepository(ctx, client.Owner, client.Repo)
	if err != nil {
		log.Fatal(err)
	}

	db.InsertRepo(ctx, repo.ID, repo.Name, repo.URL)

	issues, err := client.ListIssues(ctx, client.Owner, client.Repo)
	if err != nil {
		log.Fatal(err)
	}

	for _, issue := range issues {
		db.InsertIssue(ctx, issue.ID, repo.ID, issue.Title, issue.Number, issue.CreatedAt, issue.ClosedAt)
	}
}
