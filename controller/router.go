package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/yanrishbe/booking-service/model"
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

	// api.HandleFunc("/login", login).Methods(http.MethodPost)
	// api.HandleFunc("/protected", ProtectedEndpoint).Methods(http.MethodGet)
	// api.HandleFunc("/test", ValidateMiddleware(TestEndpoint)).Methods(http.MethodGet)
	return api
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if err != nil {
					log.Fatal(json.NewEncoder(w).Encode(model.Exception{Message: err.Error()}))
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					log.Fatal(json.NewEncoder(w).Encode(model.Exception{Message: "Invalid authorization token"}))
				}
			}
		} else {
			log.Fatal(json.NewEncoder(w).Encode(model.Exception{Message: "An authorization header is required"}))
		}
	})
}
