package dto

import (
	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
)

func ArrayMenuToResponse(menus []entity.Menu) presentation.Response {
	resp := presentation.Response{
		Data: menus,
	}

	return resp
}
