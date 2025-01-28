package domain

import "time"

type Subscribe struct {
	ID        string
	Email     string
	CreatedAt time.Time
}
