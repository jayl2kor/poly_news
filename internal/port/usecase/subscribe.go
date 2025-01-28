package usecase

import (
	"context"
)

type Subscribe interface {
	Subscribe(_ context.Context, email string) error
	Unsubscribe(_ context.Context, email string) error
}
