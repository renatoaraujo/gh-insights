package github

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v48/github"
)

type Issue struct {
	ID        int64
	Title     string
	Number    int
	State     string
	CreatedAt time.Time
	ClosedAt  time.Time
}

func (gh GitHub) ListIssues(ctx context.Context, owner, repo string) ([]*Issue, error) {
	opts := &github.IssueListByRepoOptions{
		State: "all",
		ListOptions: github.ListOptions{
			PerPage: 300,
		},
	}

	var issues []*Issue

	for {
		gissues, res, err := gh.Client.Issues.ListByRepo(ctx, owner, repo, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list issues from %s/%s. %w", owner, repo, err)
		}

		for _, gissue := range gissues {
			if gissue.GetPullRequestLinks() != nil {
				continue
			}

			issue := &Issue{
				ID:        gissue.GetID(),
				Title:     gissue.GetTitle(),
				Number:    gissue.GetNumber(),
				State:     gissue.GetState(),
				CreatedAt: gissue.GetCreatedAt(),
				ClosedAt:  gissue.GetClosedAt(),
			}
			issues = append(issues, issue)
		}

		if res.NextPage == 0 {
			break
		}

		opts.Page = res.NextPage
	}

	return issues, nil
}
