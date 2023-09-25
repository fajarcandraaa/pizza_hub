package routers

import (
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	"github.com/fajarcandraaa/pizza_hub/middleware"
	"github.com/fajarcandraaa/pizza_hub/usecase"
)

func orderRouter(p *PathPrefix, s *service.Service, m *middleware.MiddlewareStruct) {
	var (
		orderUseCase = usecase.NewOrderUseCase(s)
	)

	p.Order.HandleFunc("", m.BasicAuthMiddleware(orderUseCase.PlaceOrder)).Methods("POST")
}
