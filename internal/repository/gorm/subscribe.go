package gorm

import (
	"context"
	"errors"
	"net/http"
	polyNewsErrors "poly_news/internal/errors"
	"poly_news/internal/repository/gorm/model"

	"poly_news/internal/port/domain"

	"gorm.io/gorm"
)

type Subscribe struct {
}

func NewSubscribe() *Subscribe {
	return &Subscribe{}
}

func (s *Subscribe) SubscribeWithEmail(_ context.Context, tx *gorm.DB, email string) error {
	subscribeModel := model.Subscribe{
		Email: email,
	}
	result := tx.Create(&subscribeModel)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return polyNewsErrors.New(result.Error, http.StatusConflict, "")
	}

	if result.Error != nil {
		return polyNewsErrors.New(result.Error, http.StatusInternalServerError, result.Error.Error())
	}

	return nil
}

func (s *Subscribe) UnsubscribeWithEmail(_ context.Context, tx *gorm.DB, email string) error {
	result := tx.Where("email = ?", email).Delete(&model.Subscribe{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Subscribe) GetAllSubscribes(_ context.Context, tx *gorm.DB) ([]domain.Subscribe, error) {
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
