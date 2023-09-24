package repositories

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
)

type OrderRepositoryContract interface {
	CreateOrder(ctx context.Context, payload *entity.Order) error
	UpdateFalseProgress(ctx context.Context, orderId string) error
}
