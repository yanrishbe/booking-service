package controller

import (
	"log"
	"net/http"

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
	api.Router.Use(CORS)
	userRouter := newUserRouter(userService)
	accountRouter := newAccountRouter(accountService, *userRouter)
	bookingRouter := newBookingRouter(bookingService, *userRouter)

	log.Println("start")

	api.PathPrefix(usersRoute).Handler(userRouter)
	api.PathPrefix(usersRoute).Handler(accountRouter)
	api.PathPrefix(usersRoute).Handler(bookingRouter)

	return api
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
