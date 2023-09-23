package repositories

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
)

type ChefReposirotyContract interface {
	NewChef(ctx context.Context, payload entity.Chef) error
	SignIng(ctx context.Context, payload presentation.LoginRequest) (*entity.Chef, error)
}
