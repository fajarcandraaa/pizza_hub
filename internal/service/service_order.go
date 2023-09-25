package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fajarcandraaa/pizza_hub/internal/dto"
	"github.com/fajarcandraaa/pizza_hub/internal/entity"
	"github.com/fajarcandraaa/pizza_hub/internal/presentation"
	"github.com/fajarcandraaa/pizza_hub/internal/repositories"
	"github.com/go-redis/redis/v8"
)

type orderService struct {
	repo *repositories.Repository
	rds  *redis.Client
}

func NewOrderService(repo *repositories.Repository, rds *redis.Client) *orderService {
	return &orderService{
		repo: repo,
		rds:  rds,
	}
}

// NewOrder implements OrderServiceContract.
func (s *orderService) NewOrder(ctx context.Context, payload presentation.OrderRequest) error {
	var (
		duration time.Duration
		rdsKey   = fmt.Sprintf("inProgress-%s", payload.CheffInitials)
	)

	switch payload.MenuCode {
	case "pzzachse":
		duration = 3 * time.Second
	case "pzzabq":
		duration = 5 * time.Second
	default:
		return entity.ErrMenuNotExist
	}

	rdsResult, err := s.rds.Exists(ctx, rdsKey).Result()
	if err != nil {
		return err
	}
	if rdsResult == 1 {
		return entity.ErrStillInProgressTime
	}
	dataPayload := dto.OrderRequestToDatabase(payload)

	log.Print(duration)

	go func() {
		err = s.rds.Set(ctx, rdsKey, payload.MenuCode, duration).Err()
		if err != nil {
			<-ctx.Done()
			fmt.Errorf("error set order to redis")
			return
		}

		err = s.repo.Order.CreateOrder(ctx, dataPayload)
		if err != nil {
			<-ctx.Done()
			fmt.Errorf("error create order")
			return
		}
	}()

	go func() {
		err := s.repo.Order.UpdateFalseProgress(context.Background(), dataPayload.ID)
		if err != nil {
			fmt.Errorf("error update progress status order")
		}
	}()

	return nil
}
