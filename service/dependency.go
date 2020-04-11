package service

import (
	"context"

	"github.com/yanrishbe/booking-service/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
	CheckPassword(ctx context.Context, email string) (string, error)
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, email string) error
	GetUser(ctx context.Context, email string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

type AccountRepository interface {
	CreateAccount(ctx context.Context, account model.Account) (string, error)
	UpdateAccount(ctx context.Context, account model.Account) error
	GetAccount(ctx context.Context, id string) (*model.Account, error)
}

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking model.Booking) (string, error)
	UpdateBooking(ctx context.Context, booking model.Booking) error
	DeleteBooking(ctx context.Context, id string) error
	GetBooking(ctx context.Context, id string) (*model.Booking, error)
	GetAllBookings(ctx context.Context) ([]model.Booking, error)
}
