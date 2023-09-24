package service

import "github.com/fajarcandraaa/pizza_hub/internal/repositories"

type Service struct {
	ChefService ChefServiceContract
	MenuSerice MenuServiceContract
}

func ServiceChef(repo *repositories.Repository) ChefServiceContract {
	return NewChefService(repo)
}

func ServiceMenu(repo *repositories.Repository) MenuServiceContract {
	return NewMenuService(repo)
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		ChefService: ServiceChef(repo),
		MenuSerice:  ServiceMenu(repo),
	}
}