package usecase

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fajarcandraaa/pizza_hub/helpers"
	"github.com/fajarcandraaa/pizza_hub/internal/dto"
	"github.com/fajarcandraaa/pizza_hub/internal/service"
)

type MenuUseCase struct {
	service *service.Service
}

func NewMenuUseCase(service *service.Service) *MenuUseCase {
	return &MenuUseCase{
		service: service,
	}
}

func (s *MenuUseCase) GetListMenu(w http.ResponseWriter, r *http.Request) {
	var (
		responder    = helpers.NewHTTPResponse("listMenu")
		ctx          = context.Background()
		param        = r.URL.Query()
		paramSortBy  = param.Get("sortby")
		paramOrderBy = param.Get("orderby")
		paramPerPage = param.Get("perpage")
		paramPage    = param.Get("page")
	)

	paginationParam, err := helpers.SetDefaultPginationParam(paramPage, paramPerPage, paramOrderBy, paramSortBy)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, "value of query parameters has diferent type")
		return
	}

	sortBy := paginationParam.SortBy
	orderBy := paginationParam.OrderBy
	perPage := paginationParam.PerPage
	page, _ := strconv.Atoi(paginationParam.Page)

	payload := dto.RequestParamToMeta(sortBy, orderBy, int(perPage), page)

	menuData, total, err := s.service.MenuService.List(ctx, payload)
	if err != nil {
		responder.ErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	pagination, err := helpers.GetPagination(helpers.PaginationParams{
		Path:        "list.menu",
		Page:        strconv.Itoa(page),
		TotalRows:   int32(total),
		PerPage:     int32(perPage),
		OrderBy:     orderBy,
		SortBy:      sortBy,
		CurrentPage: int32(page),
	})
	if err != nil {
		responder.ErrorJSON(w, http.StatusConflict, "error pagination")
		return
	}

	menuData.Status = "OK"
	responder.SuccessWithMeta(w, menuData, pagination, http.StatusOK, "menu list")
	return
}
