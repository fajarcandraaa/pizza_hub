package repositories

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/jinzhu/gorm"
)

type OrderRepositories struct {
	db *gorm.DB
}

func NewOrderRepositories(db *gorm.DB) *OrderRepositories {
	return &OrderRepositories{
		db: db,
	}
}

// CreateOrder implements OrderRepositoryContract.
func (r *OrderRepositories) CreateOrder(ctx context.Context, payload *entity.Order) error {
	var (
		query = `
			INSERT INTO orders (id, menu_code, cheff_initials, in_progress)
			VALUES ($1, $2, $3, $4);
		`
	)

	arg := []interface{}{
		&payload.ID,
		&payload.MenuCode,
		&payload.CheffInitials,
		&payload.InProgress,
	}

	err := r.db.Exec(query, arg...).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateProgress implements OrderRepositoryContract.
func (r *OrderRepositories) UpdateFalseProgress(ctx context.Context, orderId string) error {
	var (
		query = `
			UPDATE orders SET in_progress = false WHERE id = $1;
		`
	)

	err := r.db.Exec(query, orderId).Error
	if err != nil {
		return err
	}

	return nil
}
