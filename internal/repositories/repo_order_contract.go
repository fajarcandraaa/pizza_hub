package repositories

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
)

type OrderRepositoryContract interface {
	CreateOrder(ctx context.Context, payload *entity.Order) error
	GetOrders(ctx context.Context, meta presentation.MetaPagination) ([]entity.Order, int64, error)
	UpdateFalseProgress(ctx context.Context, orderId string) error
}