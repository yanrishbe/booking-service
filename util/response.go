package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yanrishbe/booking-service/model"
)

type GetAllUsersResponse struct {
	ID               string `json:"id,omitempty" bson:"_id,omitempty"`
	AccountID        string `json:"accountId,omitempty" bson:"accountId"`
	*BookingResponse `json:"booking,omitempty"`
	Name             string `json:"name" bson:"name"`
	Surname          string `json:"surname" bson:"surname"`
	Patronymic       string `json:"patronymic" bson:"patronymic"`
	Phone            string `json:"phone" bson:"phone" `
	Email            string `json:"email" bson:"email"`
}

func AllUsersResponse(users []model.User, bookings []model.Booking) []GetAllUsersResponse {
	var resp []GetAllUsersResponse
	for i := range users {
		user := GetAllUsersResponse{
			ID:         users[i].ID,
			AccountID:  users[i].AccountID,
			Name:       users[i].Name,
			Surname:    users[i].Surname,
			Patronymic: users[i].Patronymic,
			Phone:      users[i].Phone,
			Email:      users[i].Email,
		}
		matchBooking(&user, bookings)
		resp = append(resp, user)
	}
	return resp
}

func matchBooking(user *GetAllUsersResponse, bookings []model.Booking) {
	for i := range bookings {
		if bookings[i].UserID != nil {
			if user.ID == *bookings[i].UserID {
				user.BookingResponse = NewBookingResponse(bookings[i])
				return
			}
		}
	}
}

type GetAllBookingsResponse struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Vip     bool   `json:"isVip" bson:"vip"`
	Price   string `json:"price" bson:"price"`
	Stars   int    `json:"stars" bson:"stars"`
	Persons int    `json:"persons" bson:"persons"`
	URL     string `json:"url" bson:"url"`
	Empty   bool   `json:"empty" bson:"empty"`
}

func AllBookingsResponse(bookings []model.Booking) []GetAllBookingsResponse {
	var resp []GetAllBookingsResponse
	for i := range bookings {
		respBooking := GetAllBookingsResponse{
			ID:      bookings[i].ID,
			Vip:     bookings[i].Vip,
			Stars:   bookings[i].Stars,
			Persons: bookings[i].Persons,
			URL:     bookings[i].URL,
			Empty:   bookings[i].Empty,
		}
		cents := bookings[i].Price % 100
		rest := bookings[i].Price / 100
		price := fmt.Sprintf("%d.%d", rest, cents)
		respBooking.Price = price

		resp = append(resp, respBooking)
	}
	return resp
}

type BookingResponse struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	Vip        bool   `json:"vip" bson:"vip"`
	Price      string `json:"price" bson:"price"`
	Stars      int    `json:"stars" bson:"stars"`
	Persons    int    `json:"persons" bson:"persons"`
	Expiration string `json:"expiration" bson:"expiration"`
	MaxDays    int    `json:"maxDays" bson:"maxDays"`
}

func NewBookingResponse(booking model.Booking) *BookingResponse {
	respBooking := BookingResponse{
		ID:         booking.ID,
		Vip:        booking.Vip,
		Stars:      booking.Stars,
		Persons:    booking.Persons,
		Expiration: booking.Expiration.In(time.Local).Format(time.ANSIC),
		MaxDays:    booking.MaxDays,
	}
	cents := booking.Price % 100
	rest := booking.Price / 100
	price := fmt.Sprintf("%d.%d", rest, cents)
	respBooking.Price = price

	return &respBooking
}

type UserResponse struct {
	*AccountResponse `json:"account,omitempty"`
	*BookingResponse `json:"booking,omitempty"`
	ID               string `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string `json:"name" bson:"name"`
	Surname          string `json:"surname" bson:"surname"`
	Patronymic       string `json:"patronymic" bson:"patronymic"`
	Phone            string `json:"phone" bson:"phone" `
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
}

func NewUserResponse(user *model.User, account *model.Account, booking *model.Booking) (*UserResponse, error) {
	if user == nil {
		return nil, fmt.Errorf("could not create user response: user is empty")
	}
	userResp := UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Phone:      user.Phone,
		Email:      user.Email,
		Password:   user.Password,
	}
	var accResp *AccountResponse
	if account != nil {
		accResp = NewAccountResponse(*account)
	}
	var bookResp *BookingResponse
	if booking != nil {
		bookResp = NewBookingResponse(*booking)
	}
	userResp.AccountResponse = accResp
	userResp.BookingResponse = bookResp
	return &userResp, nil
}

type AccountResponse struct {
	ID          string `json:"id,omitempty"`
	UserID      string `json:"userId" bson:"userId"`
	CreditCard  bool   `json:"creditCard" bson:"creditCard"`
	LegalEntity bool   `json:"legalEntity" bson:"legalEntity"`
	Bank        string `json:"bank"`
	Amount      string `json:"amount"`
}

func NewAccountResponse(account model.Account) *AccountResponse {
	accResp := AccountResponse{
		ID:          account.ID,
		UserID:      account.UserID,
		CreditCard:  account.CreditCard,
		LegalEntity: account.LegalEntity,
		Bank:        account.Bank,
	}
	cents := account.Amount % 100
	rest := account.Amount / 100
	amount := fmt.Sprintf("%d.%d", rest, cents)
	accResp.Amount = amount
	return &accResp
}

type TokenDetails struct {
	UserID            string `json:"userId"`
	AccessToken       string `json:"token"`
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
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		log.Println("encoding error")
		return
	}
}

func JSONError(code int, w http.ResponseWriter, err error) {
	w.WriteHeader(code)
	err = json.NewEncoder(w).Encode(ErrorMessage{
		Error: err.Error(),
	})
	if err != nil {
		log.Println("encoding error")
	}
}
