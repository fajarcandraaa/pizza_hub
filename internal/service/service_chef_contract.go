package service

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
)

type ChefServiceContract interface {
	Insert(ctx context.Context, payload presentation.NewChefRequest) error
	Login(ctx context.Context, username, password string) (*presentation.Response, error)
}
