package routers

import (
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	"github.com/fajarcandraaa/pizza_hub/usecase"
)

func chefRouter(p *PathPrefix, s *service.Service) {
	var (
		chefUseCase = usecase.NewChefUseCase(s)
	)

	p.Chef.HandleFunc("/login", chefUseCase.Authentication).Methods("POST")
}
