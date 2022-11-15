package github

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v48/github"
)

type Pull struct {
	ID        int64
	Title     string
	Number    int
	State     string
	CreatedAt time.Time
	ClosedAt  time.Time
}

func (gh GitHub) ListPulls(ctx context.Context, owner, repo string) ([]*Pull, error) {
	opts := &github.PullRequestListOptions{
		State: "all",
		ListOptions: github.ListOptions{
			PerPage: 300,
		},
	}

	var pulls []*Pull

	for {
		prs, res, err := gh.Client.PullRequests.List(ctx, owner, repo, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list pulls from %s/%s. %w", owner, repo, err)
		}

		for _, pr := range prs {
			pull := &Pull{
				ID:        pr.GetID(),
				Title:     pr.GetTitle(),
				Number:    pr.GetNumber(),
				State:     pr.GetState(),
				CreatedAt: pr.GetCreatedAt(),
				ClosedAt:  pr.GetClosedAt(),
			}

			pulls = append(pulls, pull)
		}

		if res.NextPage == 0 {
			break
		}

		opts.Page = res.NextPage
	}

	return pulls, nil
}
