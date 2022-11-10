package github

import (
	"context"
	"fmt"
)

type Repository struct {
	ID   int64
	Name string
	URL  string
}

func (gh GitHub) GetRepository(ctx context.Context, owner, name string) (*Repository, error) {
	repo, _, err := gh.Client.Repositories.Get(ctx, owner, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get repository %s/%s. %w", owner, name, err)
	}

	repository := &Repository{
		ID:   repo.GetID(),
		Name: repo.GetName(),
		URL:  repo.GetURL(),
	}

	return repository, nil
}
