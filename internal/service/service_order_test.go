package service_test

import (
	"context"
	"testing"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	"github.com/fajarcandraaa/pizza_hub/test/faker"
	"github.com/stretchr/testify/require"
)

func TestPlaceOrder(t *testing.T) {
	db, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Order{})

	rds, err := testRedisConfig(t)
	require.NoError(t, err)

	r := repositories.NewRepository(db)
	s := service.NewOrderService(r, rds)

	t.Run("Order: if url is valid, expected no error", func(t *testing.T) {
		var (
			ctx = context.Background()
		)
		payload := presentation.OrderRequest{
			MenuCode:      "pzzabq",
			CheffInitials: faker.FakeInitialsChef01,
		}

		err := s.NewOrder(ctx, payload)
		require.NoError(t, err)
		require.Equal(t, err, nil)
	})

	t.Run("Order: if url is not valid, expected error", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		payload := presentation.OrderRequest{}

		err := s.NewOrder(ctx, payload)
		require.Error(t, err)
	})
}
