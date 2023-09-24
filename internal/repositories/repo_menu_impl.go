package repositories

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/jinzhu/gorm"
)

type MenuRepositories struct {
	db *gorm.DB
}

func NewMenuRepositories(db *gorm.DB) *MenuRepositories {
	return &MenuRepositories{
		db: db,
	}
}

// GetMenus implements MenuRepositoryContract.
func (r *MenuRepositories) GetMenus(ctx context.Context, meta presentation.MetaPagination) ([]entity.Menu, int64, error) {
	var (
		count int64
		menus = make([]entity.Menu, 0)

		offset = (meta.Page - 1) * meta.PerPage
		order  = meta.OrderBy + " " + meta.SortBy
		model  = r.db.Model(&entity.Menu{})
	)

	if err := model.Order(order).Limit(meta.PerPage).Offset(offset).Find(&menus).Error; err != nil {
		return nil, 0, err
	}

	_ = model.Count(&count)

	return menus, count, nil
}
