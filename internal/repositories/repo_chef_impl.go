package repositories

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/jinzhu/gorm"
)

type ChefRepositories struct {
	db *gorm.DB
}

func NewChefRepositories(db *gorm.DB) *ChefRepositories {
	return &ChefRepositories{
		db: db,
	}
}

// SignIng implements ChefReposirotyContract.
func (r *ChefRepositories) SignIng(ctx context.Context, payload presentation.LoginRequest) (*entity.Chef, error) {
	var (
		chef = entity.Chef{}
	)

	err := r.db.Model(entity.Chef{}).Where("username = ?", payload.Username).Take(&chef).Error
	if err != nil {
		return nil, err
	}

	return &chef, nil
}

// NewChef implements ChefReposirotyContract.
func (r *ChefRepositories) NewChef(ctx context.Context, payload entity.Chef) error {
	panic("unimplemented")
}
