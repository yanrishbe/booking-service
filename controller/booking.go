package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yanrishbe/booking-service/util"
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

	router.Path(bookingsRoute).Methods(http.MethodPost).HandlerFunc(validateTokenMiddleware(router.createBooking))
	router.Path(bookingsRoute).Methods(http.MethodGet).HandlerFunc(validateTokenMiddleware(router.getAllBookings))
	router.Path(bookingsRoute + "/{bookingId}").Methods(http.MethodGet).HandlerFunc(validateTokenMiddleware(router.getBooking))
	router.Path(bookingsRoute + "/{bookingId}").Methods(http.MethodPut).HandlerFunc(validateTokenMiddleware(router.updateBooking))
	router.Path(bookingsRoute + "/{bookingId}").Methods(http.MethodDelete).HandlerFunc(validateTokenMiddleware(router.deleteBooking))

	return &router
}

func (br bookingRouter) createBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}

	var b util.BookingRequest
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.JSONError(http.StatusUnprocessableEntity, w, err)
		return
	}
	err = br.service.Create(r.Context(), b, id)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (br bookingRouter) getBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}
	bookingId := mux.Vars(r)["bookingId"]
	response, err := br.service.Get(r.Context(), bookingId)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	util.JSON(w, response)
}

func (br bookingRouter) getAllBookings(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}
	response, err := br.service.GetAll(r.Context())
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	util.JSON(w, response)
}

func (br bookingRouter) updateBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}

	var b util.BookingRequest
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.JSONError(http.StatusUnprocessableEntity, w, err)
		return
	}

	bookingID := mux.Vars(r)["bookingId"]
	err = br.service.Update(r.Context(), b, bookingID, id)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (br bookingRouter) deleteBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}
	bookingId := mux.Vars(r)["bookingId"]
	err = br.service.Delete(r.Context(), bookingId, id)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
