package dto

import (
	"time"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/google/uuid"
)

func ChefRequestToDatabase(payload presentation.NewChefRequest) entity.Chef {
	resp := entity.Chef{
		ID:            uuid.NewString(),
		CheffInitials: payload.CheffInitials,
		CheffName:     payload.CheffName,
		Username:      payload.Username,
		Password:      payload.Password,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}

	return resp
}
