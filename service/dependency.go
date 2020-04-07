package service

import (
	"context"

	"github.com/yanrishbe/booking-service/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) (string, error)
}

type AccountRepository interface {
}

type BookingRepository interface {
}
