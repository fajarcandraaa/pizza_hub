package usecase

import (
	"net/http"

	"github.com/fajarcandraaa/pizza_hub/internal/service"
)

type ChefUseCase struct {
	service *service.Service
}

func NewChefUseCase(service *service.Service) *ChefUseCase {
	return &ChefUseCase{
		service: service,
	}
}

func (u *ChefUseCase) Authentication(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (u *ChefUseCase) AddNewCheff(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
