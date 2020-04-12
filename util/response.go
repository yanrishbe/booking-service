package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yanrishbe/booking-service/model"
)

type UserResponse struct {
	*model.Account `json:"account,omitempty"`
	*model.Booking `json:"booking,omitempty"`
	ID             string `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string `json:"name" bson:"name"`
	Surname        string `json:"surname" bson:"surname"`
	Patronymic     string `json:"patronymic" bson:"patronymic"`
	Phone          string `json:"phone" bson:"phone" `
	Email          string `json:"email" bson:"email"`
	Password       string `json:"password" bson:"password"`
}

func NewUserResponse(user *model.User, account *model.Account, booking *model.Booking) (*UserResponse, error) {
	if user == nil {
		return nil, fmt.Errorf("could not create user response: user is empty")
	}
	return &UserResponse{
		Account:    account,
		Booking:    booking,
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Phone:      user.Phone,
		Email:      user.Email,
		Password:   user.Password,
	}, nil
}

type TokenDetails struct {
	AccessToken       string `json:"access_token"`
	Role              string `json:"role"`
	AccessExpiration  int64  `json:"access_expiration"`
	AccessUuid        string `json:"access_uuid,omitempty"`
	RefreshExpiration int64  `json:"refresh_expiration,omitempty"`
	RefreshToken      string `json:"refresh_token,omitempty"`
	RefreshUuid       string `json:"refresh_uuid,omitempty"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		log.Println("encoding error")
		return
	}
}

func JSONError(code int, w http.ResponseWriter, err error) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(ErrorMessage{
		Error: err.Error(),
	})
	if err != nil {
		log.Println("encoding error")
	}
}
