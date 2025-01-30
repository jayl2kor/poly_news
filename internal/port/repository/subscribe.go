package repository

import (
	"context"
	"poly_news/internal/port/domain"

	"gorm.io/gorm"
)

type Subscribe interface {
	SubscribeWithEmail(_ context.Context, tx *gorm.DB, email string) error
	UnsubscribeWithEmail(_ context.Context, tx *gorm.DB, email string) error
	GetAllSubscribes(_ context.Context, tx *gorm.DB) ([]domain.Subscribe, error)

	//CreateSubscribeHistory(_ context.Context, tx gorm.DB, subscribe domain.Subscribe, flag bool) error
}
