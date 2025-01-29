package usecase

import (
	"context"
	"poly_news/internal/port/domain"
	"poly_news/internal/port/repository"

	"gorm.io/gorm"
)

type Subscribe struct {
	conn                *gorm.DB
	subscribeRepository repository.Subscribe
}

func NewSubscribe(conn *gorm.DB, subscribeRepository repository.Subscribe) *Subscribe {
	return &Subscribe{
		conn:                conn,
		subscribeRepository: subscribeRepository,
	}
}

func (s *Subscribe) SubscribeWithEmail(ctx context.Context, email string) error {
	return nil
}

func (s *Subscribe) UnsubscribeWithEmail(ctx context.Context, email string) error {
	return nil
}

func (s *Subscribe) GetAllSubscribes(ctx context.Context) ([]domain.Subscribe, error) {
	return nil, nil
}
