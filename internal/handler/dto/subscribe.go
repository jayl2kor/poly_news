package dto

import "poly_news/internal/port/domain"

type SubscribeRequest struct {
	Email string `json:"email"`
}

type UnsubscribeRequest struct {
	Email string `json:"email"`
}

type ListSubscribeResponse struct {
	Subscribes []domain.Subscribe `json:"subscribes"`
}
