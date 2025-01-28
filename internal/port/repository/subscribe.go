package repository

import (
	"context"
	"poly_news/internal/port/domain"

	"gorm.io/gorm"
)

type Subscribe interface {
	Subscribe(_ context.Context, tx gorm.DB, subscribe domain.Subscribe) error
	Unsubscribe(_ context.Context, tx gorm.DB, subscribe domain.Subscribe) error

	CreateSubscribeHistory(_ context.Context, tx gorm.DB, subscribe domain.Subscribe, flag bool) error
}
