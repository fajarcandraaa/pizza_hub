package service

import (
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	ChefService  ChefServiceContract
	MenuService  MenuServiceContract
	OrderService OrderServiceContract
}

func ServiceChef(repo *repositories.Repository) ChefServiceContract {
	return NewChefService(repo)
}

func ServiceMenu(repo *repositories.Repository) MenuServiceContract {
	return NewMenuService(repo)
}

func ServiceOrder(repo *repositories.Repository, rds *redis.Client) OrderServiceContract {
	return NewOrderService(repo, rds)
}

func NewService(repo *repositories.Repository, rds *redis.Client) *Service {
	return &Service{
		ChefService:  ServiceChef(repo),
		MenuService:  ServiceMenu(repo),
		OrderService: ServiceOrder(repo, rds),
	}
}
