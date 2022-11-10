package github

import "github.com/google/go-github/v48/github"

type GitHub struct {
	Client *github.Client
	Owner  string
	Repo   string
}

func NewPublic(owner, repo string) GitHub {
	return GitHub{
		Client: github.NewClient(nil),
		Owner:  owner,
		Repo:   repo,
	}
}
