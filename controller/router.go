package controller

import (
	"fmt"

	"github.com/gorilla/mux"
)

type API struct {
	*mux.Router
}

const (
	userRoute    = "/user"
	accountRoute = "/account"
	bookingRoute = "/booking"
)

// NewRouter creates a router for booking-service API.
func NewRouter(userService User, accountService Account, bookingService Booking) API {
	api := API{
		Router: mux.NewRouter(),
	}
	userRouter := newUserRouter(userService)
	accountRouter := newAccountRouter(accountService)
	bookingRouter := newBookingRouter(bookingService)

	fmt.Println("Starting the application...")

	api.PathPrefix(userRoute).Handler(userRouter)
	api.PathPrefix(accountRoute).Handler(accountRouter)
	api.PathPrefix(bookingRoute).Handler(bookingRouter)

	return api
}
