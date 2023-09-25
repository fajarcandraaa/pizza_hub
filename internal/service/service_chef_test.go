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

func TestAddNewChef(t *testing.T) {
	db, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Chef{})

	r := repositories.NewRepository(db)
	s := service.NewChefService(r)

	t.Run("Chef: if url is valid, expected no error", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		payload := presentation.NewChefRequest{
			CheffInitials: faker.FakeInitialsChef01,
			CheffName:     faker.FakeNameChef01,
		}

		err := s.Insert(ctx, payload)
		require.NoError(t, err)
		require.Equal(t, err, nil)
	})

	t.Run("Chef: if url is not valid, expected error", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		payload := presentation.NewChefRequest{
			CheffInitials: faker.FakeInitialsChef01,
		}

		err := s.Insert(ctx, payload)
		require.Error(t, err)
	})
}
