package github

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v48/github"
)

type Issue struct {
	ID        int64
	CreatedAt time.Time
	ClosedAt  time.Time
}

func (gh GitHub) ListIssues(ctx context.Context, owner, repo string) ([]*Issue, error) {
	opts := &github.IssueListByRepoOptions{
		State: "closed",
	}
	gissues, _, err := gh.Client.Issues.ListByRepo(ctx, owner, repo, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list issues from %s/%s", owner, repo)
	}

	var issues []*Issue
	for _, gissue := range gissues {
		issue := &Issue{
			ID:        gissue.GetID(),
			CreatedAt: gissue.GetCreatedAt(),
			ClosedAt:  gissue.GetClosedAt(),
		}
		issues = append(issues, issue)
	}

	return issues, nil
}
