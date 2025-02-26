package controller

import (
	"context"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

// User defines interface for user-related CRUD operations.
type User interface {
	Create(ctx context.Context, user model.User) (string, error)
	Login(ctx context.Context, loginRequest util.LoginRequest) (string, error)
	Get(ctx context.Context, id string) (*util.UserResponse, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, userRequest util.UpdateUserRequest) error
	GetAllUsers(ctx context.Context) ([]util.GetAllUsersResponse, error)
}

// Booking defines interface for booking-related CRUD operations.
type Booking interface {
	Create(ctx context.Context, newBookingReq util.BookingRequest, userID string) error
	Update(ctx context.Context, newBookingReq util.BookingRequest, bookingID string, userID string) error
	Get(ctx context.Context, id string) (*util.BookingResponse, error)
	GetAll(ctx context.Context) ([]util.GetAllBookingsResponse, error)
	Delete(ctx context.Context, bookingID string, userID string) error
}

// Account defines interface for account-related CRU operations.
type Account interface {
	Create(ctx context.Context, accountRequest util.AccountRequest, userID string) (string, error)
	Get(ctx context.Context, id string) (*util.AccountResponse, error)
	Update(ctx context.Context, newAccountReq util.AccountRequest, accountID string, userID string) error
}
