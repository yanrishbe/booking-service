package controller

import (
	"log"

	"github.com/gorilla/mux"
)

type API struct {
	*mux.Router
}

const usersRoute = "/users"

// NewRouter creates a router for booking-service API.
func NewRouter(userService User, accountService Account, bookingService Booking) API {
	api := API{
		Router: mux.NewRouter(),
	}
	userRouter := newUserRouter(userService)
	accountRouter := newAccountRouter(accountService, *userRouter)
	bookingRouter := newBookingRouter(bookingService, *userRouter)

	log.Println("start")

	api.PathPrefix(usersRoute).Handler(userRouter)
	api.PathPrefix(usersRoute).Handler(accountRouter)
	api.PathPrefix("bookingsRoute").Handler(bookingRouter)

	return api
}
