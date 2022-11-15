package github

import (
	"context"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type GitHub struct {
	Client *github.Client
	Owner  string
	Repo   string
}

// NewPublicClient will open a connection with GitHub API as an unauthenticated user (rate limit for requests 60/h)
func NewPublicClient(owner, repo string) GitHub {
	return GitHub{
		Client: github.NewClient(nil),
		Owner:  owner,
		Repo:   repo,
	}
}

// NewAuthenticatedClient will open a connection with GitHub API as an authenticated user (rate limit for requests 5000/h)
func NewAuthenticatedClient(owner, repo string, token string) GitHub {
	ctx := context.Background()
	tokeSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, tokeSource)

	return GitHub{
		Client: github.NewClient(httpClient),
		Owner:  owner,
		Repo:   repo,
	}
}
