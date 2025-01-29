package usecase

import (
	"context"
	"poly_news/internal/port/domain"
)

type Subscribe interface {
	SubscribeWithEmail(_ context.Context, email string) error
	UnsubscribeWithEmail(_ context.Context, email string) error

	GetAllSubscribes(_ context.Context) ([]domain.Subscribe, error)
}
