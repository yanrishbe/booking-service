package controller

import (
	"context"

	"github.com/yanrishbe/booking-service/model"
)

// User defines interface for user-related CRUD operations.
type User interface {
	Create(ctx context.Context, user model.User) (string, error)
	// Login ()
	// Edit ()
	// Get ()
	// GetAll()
	// Delete()
	// DeleteAll()
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
	// Create ()
	// Edit ()
	// Get ()
}
