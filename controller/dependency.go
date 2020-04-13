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
}

// Booking defines interface for booking-related CRUD operations.
type Booking interface {
	// Create ()
	// Edit ()
	// Get ()
	// GetAll()
	// Delete()
	// DeleteAll()
}

// Account defines interface for account-related CRU operations.
type Account interface {
	Create(ctx context.Context, account model.Account, userID string) (string, error)
	Get(ctx context.Context, id string) (*model.Account, error)
	// Edit ()
}
