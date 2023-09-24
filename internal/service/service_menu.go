package service

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/dto"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
)

type menuService struct {
	repo *repositories.Repository
}

func NewMenuService(repo *repositories.Repository) *menuService {
	return &menuService{
		repo: repo,
	}
}

// List implements MenuServiceContract.
func (s *menuService) List(ctx context.Context, meta presentation.MetaPagination) (*presentation.Response, int64, error) {
	menuList, total, err := s.repo.Menu.GetMenus(ctx, meta)
	if err != nil {
		return nil, 0, err
	}

	resp := dto.ArrayMenuToResponse(menuList)
	return &resp, total, nil
}
