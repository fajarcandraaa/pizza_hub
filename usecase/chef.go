package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fajarcandraaa/pizza_hub/helpers"
	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/fajarcandraaa/pizza_hub/internal/service"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type ChefUseCase struct {
	service *service.Service
}

func NewChefUseCase(service *service.Service) *ChefUseCase {
	return &ChefUseCase{
		service: service,
	}
}

func (u *ChefUseCase) Authentication(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("loginUser")
		ctx       = context.Background()
		payload   presentation.LoginRequest
	)
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
		validation.Field(&payload.Username, validation.Required),
		validation.Field(&payload.Password, validation.Required),
	)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	login, err := u.service.ChefService.Login(ctx, payload)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrUserAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		case entity.ErrPermissionNotAllowed:
			responder.FieldErrors(w, err, http.StatusBadRequest, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}

	login.Status = "OK"

	responder.SuccessJSON(w, login, http.StatusOK, "Login success")
	return
}

func (u *ChefUseCase) AddNewCheff(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
