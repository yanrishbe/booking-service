package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/yanrishbe/booking-service/util"
)

func createToken(email string) (*util.TokenDetails, error) {
	token := util.TokenDetails{
		AccessExpiration: time.Now().Add(time.Minute * 3).Unix(),
	}
	atClaims := jwt.MapClaims{
		"authorized": true,
		"email":      email,
		"expiration": token.AccessExpiration,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	var err error
	token.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, fmt.Errorf("could not create token %v", err)
	}
	return &token, nil
	/*td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["email"] = email
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	var err error
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["email"] = email
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil*/
}

// func extractToken(r *http.Request) string {
// 	authorizationHeader := r.Header.Get("Authorization")
// 	if authorizationHeader != "" {
// 		bearerToken := strings.Split(authorizationHeader, " ")
// 		if len(bearerToken) == 2 {
// 			return bearerToken[1]
// 		}
// 	}
// 	return ""
// }

func verifyToken(r *http.Request) (*jwt.Token, error) {
	// tokenString := extractToken(r)
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, fmt.Errorf("an authorization header is required")
	}
	// func in params return s key to validate a token
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
		_, ok := token.Claims.(jwt.Claims)
		if !ok {
			util.JSONError(http.StatusUnauthorized, w, fmt.Errorf("token payload is invalid"))
			return
		}
		if !token.Valid {
			util.JSONError(http.StatusUnauthorized, w, fmt.Errorf("token is invalid"))
			return
		}
		next(w, r)
	}
}
