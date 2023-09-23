package dto

import "github.com/fajarcandraaa/pizza_hub/internal/presentation"

func Login(username, password string) presentation.LoginRequest {
	resp := presentation.LoginRequest{
		Username: username,
		Password: password,
	}

	return resp
}

func ToResponse(data interface{}) presentation.Response {
	res := presentation.Response{
		Data: data,
	}

	return res
}

func RequestParamToMeta(sortBy, orderBy string, perPage, page int) presentation.MetaPagination {
	resp := presentation.MetaPagination{
		SortBy:  sortBy,
		OrderBy: orderBy,
		PerPage: perPage,
		Page:    page,
	}

	return resp
}
