package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type bookingRouter struct {
	*mux.Router
	service Booking
}

// todo DONT FORGET TO CREATE CORRECT USERS SUBROUTER AND ADD MIDDLEWARE
func newBookingRouter(service Booking) *bookingRouter {
	router := bookingRouter{
		mux.NewRouter().PathPrefix("bookingsRoute").Subrouter(),
		service,
	}

	router.Path("").Methods(http.MethodPost).HandlerFunc(router.createBooking)
	router.Path("/{id}").Methods(http.MethodGet).HandlerFunc(router.getBooking)
	router.Path("/{id}").Methods(http.MethodPost).HandlerFunc(router.updateBooking)
	router.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(router.deleteBooking)

	return &router
}

func (br bookingRouter) createBooking(w http.ResponseWriter, r *http.Request) {
}
func (br bookingRouter) getBooking(w http.ResponseWriter, r *http.Request) {
}
func (br bookingRouter) updateBooking(w http.ResponseWriter, r *http.Request) {
}
func (br bookingRouter) deleteBooking(w http.ResponseWriter, r *http.Request) {
}
