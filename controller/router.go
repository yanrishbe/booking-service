package controller

import (
	"encoding/json"
	"fmt"
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

func New() API {
	api := API{
		Router: mux.NewRouter(),
	}
	fmt.Println("Starting the application...")

	api.HandleFunc("/authenticate", CreateTokenEndpoint).Methods("POST")
	api.HandleFunc("/protected", ProtectedEndpoint).Methods("GET")
	api.HandleFunc("/test", ValidateMiddleware(TestEndpoint)).Methods("GET")
	return api
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(model.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(model.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: "An authorization header is required"})
		}
	})
}
