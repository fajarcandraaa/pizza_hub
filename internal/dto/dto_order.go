package dto

import (
	"time"

	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/google/uuid"
)

func OrderRequestToDatabase(payload presentation.OrderRequest) *entity.Order {
	resp := &entity.Order{
		ID:            uuid.NewString(),
		MenuCode:      payload.MenuCode,
		CheffInitials: payload.CheffInitials,
		InProgress:    true,
	}

	return resp
}

func TimeSleepDependMenu(menucode string) {
	if menucode == "pzzachse" { // Pizza Cheese
		time.Sleep(3 * time.Second)
	}

	if menucode == "pzzabq" { // Pizza BBQ
		time.Sleep(5 * time.Second)
	}

	return
}
