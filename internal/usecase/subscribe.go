package usecase

import (
	"context"
	"gorm.io/gorm"
	polyNewsErrors "poly_news/internal/errors"
	"poly_news/internal/port/domain"
	"poly_news/internal/port/repository"
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
	err := s.conn.Transaction(func(tx *gorm.DB) error {
		err := s.subscribeRepository.SubscribeWithEmail(ctx, tx, email)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return polyNewsErrors.Wrap(err)
	}

	return nil
}

func (s *Subscribe) UnsubscribeWithEmail(ctx context.Context, email string) error {
	err := s.conn.Transaction(func(tx *gorm.DB) error {
		err := s.subscribeRepository.UnsubscribeWithEmail(ctx, s.conn, email)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return polyNewsErrors.Wrap(err)
	}

	return nil
}

func (s *Subscribe) GetAllSubscribes(ctx context.Context) ([]domain.Subscribe, error) {
	subscribes, err := s.subscribeRepository.GetAllSubscribes(ctx, s.conn)
	if err != nil {
		return nil, err
	}
	return subscribes, nil
}
