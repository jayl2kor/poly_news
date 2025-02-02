package domain

import "time"

type NewsLetter struct {
	Title     string
	Text      string
	CreatedAt time.Time
}
