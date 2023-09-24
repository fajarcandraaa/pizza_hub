package routers

import (
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
	"github.com/fajarcandraaa/pizza_hub/internal/service"
)

func (se *Serve) initializeRoutes() {
	p := RouterConfigPrefix(se)            // set grouping prefix
	r := repositories.NewRepository(se.DB) //initiate repository
	s := service.NewService(r)             //initiate service

	//initiate endpoint
	chefRouter(p, s)
	menuRouter(p, s)
}
