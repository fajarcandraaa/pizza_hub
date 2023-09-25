package service_test

import (
	"context"
	"testing"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	"github.com/stretchr/testify/require"
)

func TestListMenu(t *testing.T) {
	db, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Menu{})

	r := repositories.NewRepository(db)
	s := service.NewMenuService(r)

	t.Run("Menu: if url is valid, expected no error", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		payload := presentation.MetaPagination{
			SortBy:  "ASC",
			OrderBy: "created_at",
			PerPage: 10,
			Page:    1,
		}

		_, _, err := s.List(ctx, payload)
		require.NoError(t, err)
		require.Equal(t, err, nil)
	})

	t.Run("Menu: if url is not valid, expected error", func(t *testing.T) {
		var (
			ctx = context.Background()
		)

		payload := presentation.MetaPagination{
			SortBy:  "ASC",
			OrderBy: "name",
			PerPage: 10,
			Page:    9,
		}

		_, _, err := s.List(ctx, payload)
		require.Error(t, err)
	})
}
