package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Chef ChefReposirotyContract
	Menu MenuRepositoryContract
}

func NewChef(db *gorm.DB) ChefReposirotyContract {
	return NewChefRepositories(db)
}

func NewMenu(db *gorm.DB) MenuRepositoryContract{
	return NewMenuRepositories(db)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Chef: NewChef(db),
		Menu: NewMenu(db),
	}
}
