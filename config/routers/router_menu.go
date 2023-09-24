package routers

import (
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	"github.com/fajarcandraaa/pizza_hub/middleware"
	"github.com/fajarcandraaa/pizza_hub/usecase"
)

func menuRouter(p *PathPrefix, s *service.Service)  {
	var (
		menuUseCase = usecase.NewMenuUseCase(s)
	)

	p.Menu.HandleFunc("", middleware.Authentication(menuUseCase.GetListMenu)).Methods("GET")
}