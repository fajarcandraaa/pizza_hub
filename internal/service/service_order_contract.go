package service

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
)

type OrderServiceContract interface {
	NewOrder(ctx context.Context, payload presentation.OrderRequest) error
	ListOrder(ctx context.Context, meta presentation.MetaPagination) (*presentation.Response, error)
}
