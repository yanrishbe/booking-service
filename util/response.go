package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yanrishbe/booking-service/model"
)

type GetAllBookingsResponse struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Vip     bool   `json:"vip" bson:"vip"`
	Price   string `json:"price" bson:"price"`
	Stars   int    `json:"stars" bson:"stars"`
	Persons int    `json:"persons" bson:"persons"`
}

func NewGetAllBookingsResponse(bookings []model.Booking) []GetAllBookingsResponse {
	var resp []GetAllBookingsResponse
	for i := range bookings {
		respBooking := GetAllBookingsResponse{
			ID:      bookings[i].ID,
			Vip:     bookings[i].Vip,
			Stars:   bookings[i].Stars,
			Persons: bookings[i].Persons,
		}
		cents := bookings[i].Price % 100
		rest := bookings[i].Price / 100
		price := fmt.Sprintf("%d.%d", rest, cents)
		respBooking.Price = price

		resp = append(resp, respBooking)
	}
	return resp
}

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

type AccountResponse struct {
	ID     string `json:"id,omitempty"`
	Bank   string `json:"bank"`
	Amount string `json:"amount"`
}

func NewAccountResponse(account model.Account) *AccountResponse {
	accResp := AccountResponse{
		ID:   account.ID,
		Bank: account.Bank,
	}
	cents := account.Amount % 100
	rest := account.Amount / 100
	amount := fmt.Sprintf("%d.%d", rest, cents)
	accResp.Amount = amount
	return &accResp
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
