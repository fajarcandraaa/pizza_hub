package service

import (
	"context"

	"github.com/fajarcandraaa/pizza_hub/config/auth"
	"github.com/fajarcandraaa/pizza_hub/helpers"
	"github.com/fajarcandraaa/pizza_hub/internal/dto"
	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
	"github.com/pkg/errors"
)

type chefService struct {
	repo *repositories.Repository
}

func NewChefService(repo *repositories.Repository) *chefService {
	return &chefService{
		repo: repo,
	}
}

// Login implements ChefServiceContract.
func (s *chefService) Login(ctx context.Context, payload presentation.LoginRequest) (*presentation.Response, error) {
	signIn, err := s.repo.Chef.SignIng(ctx, payload)
	if err != nil {
		return nil, err
	}

	checkPassword := helpers.CheckPasswordHash(payload.Password, signIn.Password)
	if !checkPassword {
		return nil, entity.ErrPermissionNotAllowed
	}

	getToken, err := auth.CreateToken(signIn.CheffInitials)
	if err != nil {
		return nil, errors.Wrap(err, "build token")
	}

	token := dto.TokenToResponse(getToken)
	resp := dto.ToResponse(token)
	return &resp, nil
}

// Insert implements ChefServiceContract.
func (s *chefService) Insert(ctx context.Context, payload presentation.NewChefRequest) error {
	panic("unimplemented")
}
