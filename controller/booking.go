package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type bookingRouter struct {
	*mux.Router
	service Booking
}

const bookingsRoute = "/{id}/bookings"

func newBookingRouter(service Booking, userRouter userRouter) *bookingRouter {
	router := bookingRouter{
		userRouter.PathPrefix("").Subrouter(),
		service,
	}

	router.Path("").Methods(http.MethodPost).HandlerFunc(validateTokenMiddleware(router.createBooking))
	router.Path("/{id}").Methods(http.MethodGet).HandlerFunc(validateTokenMiddleware(router.getBooking))
	router.Path("/{id}").Methods(http.MethodPost).HandlerFunc(validateTokenMiddleware(router.updateBooking))
	router.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(validateTokenMiddleware(router.deleteBooking))

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
