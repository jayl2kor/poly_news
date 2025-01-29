package gorm

import (
	"context"
	"poly_news/internal/repository/gorm/model"

	"poly_news/internal/port/domain"

	"gorm.io/gorm"
)

type Subscribe struct {
}

func NewSubscribe() *Subscribe {
	return &Subscribe{}
}

func (s *Subscribe) Subscribe(_ context.Context, tx gorm.DB, subscribe domain.Subscribe) error {
	subscribeModel := model.Subscribe{
		ID:    subscribe.ID,
		Email: subscribe.Email,
	}
	if err := tx.Create(&subscribeModel).Error; err != nil {
		return err
	}

	return nil
}

func (s *Subscribe) Unsubscribe(_ context.Context, tx gorm.DB, subscribe domain.Subscribe) error {
	if err := tx.Where("email = ?", subscribe.Email).Delete(&model.Subscribe{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *Subscribe) GetAllSubscribes(_ context.Context, tx gorm.DB) ([]domain.Subscribe, error) {
	subscribes := []model.Subscribe{}
	if err := tx.Find(&subscribes).Error; err != nil {
		return nil, err
	}

	domainSubscribes := []domain.Subscribe{}
	for _, subscribe := range subscribes {
		domainSubscribes = append(domainSubscribes, subscribe.ToDomain())
	}

	return domainSubscribes, nil
}
