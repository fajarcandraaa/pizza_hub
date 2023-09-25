package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fajarcandraaa/pizza_hub/config/auth"
	"github.com/fajarcandraaa/pizza_hub/helpers"
	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type OrderUseCase struct {
	service *service.Service
}

func NewOrderUseCase(service *service.Service) *OrderUseCase {
	return &OrderUseCase{
		service: service,
	}
}

func (u *OrderUseCase) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("newOrder")
		ctx       = context.Background()
		payload   presentation.OrderRequest
		token     = auth.ExtractToken(r)
	)

	// token := helpers.GetToken(header)
	claims, sttsErr := helpers.ExtractClaims(token)
	if sttsErr == false {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Errorf("failed to parse claim").Error())
	}
	initialsClaim := claims["initials"]
	if v, ok := initialsClaim.(string); ok {
		initial := v
		body, err := io.ReadAll(r.Body)
		if err != nil {
			responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
			return
		}
		err = json.Unmarshal(body, &payload)
		if err != nil {
			responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
			return
		}

		err = validation.ValidateStruct(&payload,
			validation.Field(&payload.MenuCode, validation.Required),
		)
		if err != nil {
			responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
			return
		}
		payload.CheffInitials = initial

		err = u.service.OrderService.NewOrder(ctx, payload)
		if err != nil {
			causer := errors.Cause(err)
			switch causer {
			case entity.ErrStillInProgressTime:
				responder.FieldErrors(w, err, http.StatusLocked, err.Error())
				return
			case entity.ErrMenuNotExist:
				responder.FieldErrors(w, err, http.StatusNotFound, err.Error())
				return
			default:
				responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
				return
			}
		}

		resp := presentation.Response{
			Status: "OK",
		}
		responder.SuccessJSON(w, resp, http.StatusCreated, "create new oreder")
		return
	} else {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Errorf("failed to parse claim").Error())
		return
	}
}
