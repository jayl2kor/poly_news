package repository

import "poly_news/internal/domain"

type SubscribeRepository interface {
	Subscribe(subscribe domain.Subscribe)
}
