package seeder

import (
	"github.com/fajarcandraaa/pizza_hub/test/faker"
	"github.com/jinzhu/gorm"
)

func seedChef(db *gorm.DB) error {
	fakeChefData := faker.FakeChef()
	for _, data := range fakeChefData {
		err := db.Debug().Create(&data).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedMenu(db *gorm.DB) error {
	fakeMenuData := faker.FakeMenu()
	for _, data := range fakeMenuData {
		err := db.Debug().Create(&data).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func SeedData(db *gorm.DB) error {
	var count int64
	db.Table("chefs").Count(&count)
	if count == 0 {
		err := seedChef(db)
		if err != nil {
			return err
		}

		return nil
	}

	db.Table("menus").Count(&count)
	if count == 0 {
		err := seedMenu(db)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}
