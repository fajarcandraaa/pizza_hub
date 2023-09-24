package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Chef ChefReposirotyContract
}

func NewChef(db *gorm.DB) ChefReposirotyContract {
	return NewChefRepositories(db)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Chef: NewChef(db),
	}
}
