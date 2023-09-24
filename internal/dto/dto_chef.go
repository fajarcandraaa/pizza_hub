package dto

import (
	"time"

	"github.com/fajarcandraaa/pizza_hub/helpers"
	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/google/uuid"
)

// TokenToResponse is function input token as object to response
func TokenToResponse(token string) presentation.LoginResponse {
	tokenData := presentation.LoginResponse{
		Token: token,
	}

	return tokenData
}

func ChefRequestToDatabase(payload presentation.NewChefRequest) entity.Chef {
	pass, _ := helpers.HashPassword(payload.Password)
	resp := entity.Chef{
		ID:            uuid.NewString(),
		CheffInitials: payload.CheffInitials,
		CheffName:     payload.CheffName,
		Username:      payload.Username,
		Password:      pass,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}

	return resp
}
