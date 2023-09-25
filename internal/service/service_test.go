package service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	rds "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func testConfig(t *testing.T) (*gorm.DB, error) {
	err := godotenv.Load("../../.env") // Update the path accordingly
	require.NoError(t, err)
	DBDriver := os.Getenv("DB_DRIVER_TEST")
	DBHost := os.Getenv("DB_HOST_TEST")
	DBUser := os.Getenv("DB_USER_TEST")
	DBPassword := os.Getenv("DB_PASSWORD_TEST")
	DBName := os.Getenv("DB_NAME_TEST")
	DBPort := os.Getenv("DB_PORT_TEST")
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
	db, err := gorm.Open(DBDriver, DBURL) // initiate database for testing
	require.NoError(t, err)
	db.AutoMigrate(&entity.Chef{}, &entity.Menu{}, &entity.Order{})

	return db, nil

}

func testRedisConfig(t *testing.T) (*rds.Client, error) {
	// Redis connection options
	options := &rds.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Default DB
	}

	// // initiate redis client for testing
	client := rds.NewClient(options)

	return client, nil
}
