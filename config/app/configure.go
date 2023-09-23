package app

import "github.com/fajarcandraaa/pizza_hub/internal/entity"

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&entity.Chef{},
		&entity.Menu{},
		&entity.Order{},
	}

	return migrationData
}
