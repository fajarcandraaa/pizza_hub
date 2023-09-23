package service

import "github.com/fajarcandraaa/pizza_hub/internal/repositories"

type Service struct {
	ChefService ChefServiceContract
}

func ServiceChef(repo *repositories.Repository) ChefServiceContract {
	return NewChefService(repo)
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		ChefService: ServiceChef(repo),
	}
}