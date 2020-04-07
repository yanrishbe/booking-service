package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type bookingRouter struct {
	*mux.Router
	service Booking
}

func newBookingRouter(service Booking) *bookingRouter {
	router := bookingRouter{
		mux.NewRouter().PathPrefix(bookingRoute).Subrouter(),
		service,
	}

	router.Path("").Methods(http.MethodPost).HandlerFunc(router.createBooking)
	router.Path("/{id}").Methods(http.MethodGet).HandlerFunc(router.getBooking)
	router.Path("/{id}").Methods(http.MethodPost).HandlerFunc(router.updateBooking)
	router.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(router.deleteBooking)
	router.Path("").Methods(http.MethodGet).HandlerFunc(router.getAllBookings)

	return &router
}

func (br bookingRouter) createBooking(w http.ResponseWriter, r *http.Request) {
	// pID := mux.Vars(r)
	// data, err := middleware.DataFromContext(r.Context())
	// if err != nil {
	// 	middleware.JSONError(w, e.InvalidMiddlewareContext(err), http.StatusBadRequest)
	// 	return
	// }
	// err = br.service.Delete(r.Context(), data, pID["id"])
	// if err != nil {
	// 	middleware.JSONError(w, err, http.StatusInternalServerError)
	// }
	// middleware.Empty(w, http.StatusCreated)
}
func (br bookingRouter) getBooking(w http.ResponseWriter, r *http.Request) {
	// pID := mux.Vars(r)
	// data, err := middleware.DataFromContext(r.Context())
	// if err != nil {
	// 	middleware.JSONError(w, e.InvalidMiddlewareContext(err), http.StatusBadRequest)
	// 	return
	// }
	// err = ur.service.Delete(r.Context(), data, pID["id"])
	// if err != nil {
	// 	middleware.JSONError(w, err, http.StatusInternalServerError)
	// }
	// middleware.Empty(w, http.StatusCreated)
}
func (br bookingRouter) updateBooking(w http.ResponseWriter, r *http.Request) {
	// pID := mux.Vars(r)
	// data, err := middleware.DataFromContext(r.Context())
	// if err != nil {
	// 	middleware.JSONError(w, e.InvalidMiddlewareContext(err), http.StatusBadRequest)
	// 	return
	// }
	// err = ur.service.Delete(r.Context(), data, pID["id"])
	// if err != nil {
	// 	middleware.JSONError(w, err, http.StatusInternalServerError)
	// }
	// middleware.Empty(w, http.StatusCreated)
}
func (br bookingRouter) deleteBooking(w http.ResponseWriter, r *http.Request) {
	// pID := mux.Vars(r)
	// data, err := middleware.DataFromContext(r.Context())
	// if err != nil {
	// 	middleware.JSONError(w, e.InvalidMiddlewareContext(err), http.StatusBadRequest)
	// 	return
	// }
	// err = ur.service.Delete(r.Context(), data, pID["id"])
	// if err != nil {
	// 	middleware.JSONError(w, err, http.StatusInternalServerError)
	// }
	// middleware.Empty(w, http.StatusCreated)
}
func (br bookingRouter) getAllBookings(w http.ResponseWriter, r *http.Request) {
	// pID := mux.Vars(r)
	// data, err := middleware.DataFromContext(r.Context())
	// if err != nil {
	// 	middleware.JSONError(w, e.InvalidMiddlewareContext(err), http.StatusBadRequest)
	// 	return
	// }
	// err = ur.service.Delete(r.Context(), data, pID["id"])
	// if err != nil {
	// 	middleware.JSONError(w, err, http.StatusInternalServerError)
	// }
	// middleware.Empty(w, http.StatusCreated)
}
