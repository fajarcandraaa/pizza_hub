package routers

import (
	"fmt"
	"log"

	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	"github.com/fajarcandraaa/pizza_hub/middleware"
	"github.com/fajarcandraaa/pizza_hub/pkg/storage/redis"
)

func (se *Serve) initializeRoutes() {
	rds, err := redis.NewRedisConfig()
	if err != nil {
		fmt.Printf("Cannot connect to redis")
		log.Fatal("This is the error:", err)
	}

	m := middleware.MiddlewareStructFunc(se.DB)
	p := RouterConfigPrefix(se)            // set grouping prefix
	r := repositories.NewRepository(se.DB) //initiate repository
	s := service.NewService(r, rds)        //initiate service

	//initiate endpoint
	chefRouter(p, s, m)
	menuRouter(p, s, m)
	orderRouter(p, s, m)
}
