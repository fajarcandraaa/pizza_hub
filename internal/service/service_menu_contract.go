package service

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
)

type MenuServiceContract interface {
	List(ctx context.Context, meta presentation.MetaPagination) (*presentation.Response, int64, error)
}