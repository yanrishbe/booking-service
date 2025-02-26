package service

import (
	"context"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/mongo"
	"github.com/yanrishbe/booking-service/util"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
	GetPasswordAndID(ctx context.Context, email string) (*mongo.PasswordAndIDResponse, error)
	UpdateUser(ctx context.Context, userRequest util.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*model.User, error)
	UpdateAccountID(ctx context.Context, accID string, userID string) error
	UpdateBookingID(ctx context.Context, bookID string, userID string) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

type AccountRepository interface {
	CreateAccount(ctx context.Context, account model.Account) (string, error)
	GetAccount(ctx context.Context, id string) (*model.Account, error)
	UpdateAccount(ctx context.Context, account model.Account) error
	DeleteAccount(ctx context.Context, id string) error
	GetAdminAccount(ctx context.Context) (*model.Account, error)
}

type BookingRepository interface {
	GetBooking(ctx context.Context, id string) (*model.Booking, error)
	UpdateBooking(ctx context.Context, booking model.Booking) error
	GetAllBookings(ctx context.Context) ([]model.Booking, error)
}
