package service

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/dto"
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
	dataPayload := dto.ChefRequestToDatabase(payload)
	err := s.repo.Chef.NewChef(ctx, dataPayload)
	if err != nil {
		return err
	}

	return nil
}
