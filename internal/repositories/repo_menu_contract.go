package repositories

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
)

type MenuRepositoryContract interface {
	GetMenus(ctx context.Context, meta presentation.MetaPagination) ([]entity.Menu, int64, error)
}
