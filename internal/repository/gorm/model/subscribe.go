package model

import (
	"poly_news/internal/port/domain"
	"time"
)

type Subscribe struct {
	ID        int       `gorm:"primaryKey"`
	Email     string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Subscribe) TableName() string {
	return "subscribe"
}

func (s Subscribe) ToDomain() domain.Subscribe {
	return domain.Subscribe{
		ID:        s.ID,
		Email:     s.Email,
		CreatedAt: s.CreatedAt,
	}
}
