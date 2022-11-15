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
	log.Printf("Successfully saved %s/%s", client.Owner, client.Repo)

	issues, err := client.ListIssues(ctx, client.Owner, client.Repo)
	if err != nil {
		log.Fatal(err)
	}

	for _, issue := range issues {
		db.InsertIssue(ctx, issue.ID, repo.ID, issue.Title, issue.Number, issue.State, issue.CreatedAt, issue.ClosedAt)
	}
	log.Printf("Successfully saved %d issues for %s/%s", len(issues), client.Owner, client.Repo)

	pulls, err := client.ListPulls(ctx, client.Owner, client.Repo)
	if err != nil {
		log.Fatal(err)
	}

	for _, pull := range pulls {
		db.InsertPull(ctx, pull.ID, repo.ID, pull.Title, pull.Number, pull.State, pull.CreatedAt, pull.ClosedAt)
	}
	log.Printf("Successfully saved %d pulls for %s/%s", len(pulls), client.Owner, client.Repo)

	log.Printf("Finished to sync %s/%s", client.Owner, client.Repo)
}
