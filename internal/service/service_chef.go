package service

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
)

type chefService struct {
	repo *repositories.Repository
}

func NewChefService(repo *repositories.Repository) *chefService {
	return &chefService{
		repo: repo,
	}
}

// Insert implements ChefServiceContract.
func (s *chefService) Insert(ctx context.Context, payload presentation.NewChefRequest) error {
	panic("unimplemented")
}

// Login implements ChefServiceContract.
func (s *chefService) Login(ctx context.Context, username string, password string) (*presentation.Response, error) {
	panic("unimplemented")
}
