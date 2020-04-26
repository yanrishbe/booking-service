package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

func createToken(email string, id string) (*util.TokenDetails, error) {
	token := util.TokenDetails{
		AccessExpiration: time.Now().Add(time.Hour * 24).Unix(),
		UserID:           id,
	}
	var role string
	if email == model.Admin {
		role = "admin"
	} else {
		role = "user"
	}
	atClaims := jwt.MapClaims{
		"role": role,
		"id":   id,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	standardClaims := at.Claims.(jwt.MapClaims)
	standardClaims["exp"] = token.AccessExpiration
	var err error
	token.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	token.Role = role
	if err != nil {
		return nil, fmt.Errorf("could not create token %v", err)
	}
	return &token, nil
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, fmt.Errorf("an authorization header is required")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

type Authorization struct {
	Role string
	ID   string
}

func validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := verifyToken(r)
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					util.JSONError(http.StatusUnauthorized, w, fmt.Errorf("the token has expired: %v", err))
					return
				}
			}
			util.JSONError(http.StatusUnauthorized, w, err)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			util.JSONError(http.StatusUnauthorized, w, fmt.Errorf("token payload is invalid"))
			return
		}
		if !token.Valid {
			util.JSONError(http.StatusUnauthorized, w, fmt.Errorf("token is invalid"))
			return
		}
		// todo deal somehow with error here
		auth := Authorization{
			Role: claims["role"].(string),
			ID:   claims["id"].(string),
		}
		ctx := context.WithValue(r.Context(), "auth", auth)
		next(w, r.WithContext(ctx))
	}
}
