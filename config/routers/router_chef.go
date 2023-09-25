package routers

import (
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	"github.com/fajarcandraaa/pizza_hub/middleware"
	"github.com/fajarcandraaa/pizza_hub/usecase"
)

func chefRouter(p *PathPrefix, s *service.Service, m *middleware.MiddlewareStruct) {
	var (
		chefUseCase = usecase.NewChefUseCase(s)
	)

	p.Chef.HandleFunc("", m.BasicAuthMiddleware(chefUseCase.AddNewCheff)).Methods("POST")
}
