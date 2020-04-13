package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

type User struct {
	userDB    UserRepository
	accountDB AccountRepository
	bookingDB BookingRepository
}

func NewUser(userDB UserRepository, accountDB AccountRepository, bookingDB BookingRepository) *User {
	return &User{
		userDB:    userDB,
		accountDB: accountDB,
		bookingDB: bookingDB,
	}
}

func (us User) Create(ctx context.Context, user model.User) (string, error) {
	return us.userDB.CreateUser(ctx, user)
}

func (us User) Login(ctx context.Context, loginRequest util.LoginRequest) (string, error) {
	passwordAndID, err := us.userDB.GetPasswordAndID(ctx, loginRequest.Email)
	if err != nil {
		return "", err
	}
	if passwordAndID == nil {
		return "", fmt.Errorf("could not fetch users data")
	}
	return passwordAndID.ID, bcrypt.CompareHashAndPassword([]byte(passwordAndID.Password), []byte(loginRequest.Password))
}

func (us User) Get(ctx context.Context, id string) (*util.UserResponse, error) {
	user, err := us.userDB.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	var account *model.Account
	if user.AccountID != "" {
		account, err = us.accountDB.GetAccount(ctx, user.AccountID)
		if err != nil {
			return nil, err
		}
	}
	var booking *model.Booking
	if user.BookingID != "" {
		booking, err = us.bookingDB.GetBooking(ctx, user.BookingID)
		if err != nil {
			return nil, err
		}
	}
	return util.NewUserResponse(user, account, booking)
}

func (us User) Update(ctx context.Context, userRequest util.UpdateUserRequest) error {
	return us.userDB.UpdateUser(ctx, userRequest)
}

func (us User) Delete(ctx context.Context, id string) error {
	user, err := us.userDB.GetUser(ctx, id)
	if err != nil {
		return err
	}
	if user.AccountID != "" {
		err = us.accountDB.DeleteAccount(ctx, user.AccountID)
		if err != nil {
			return err
		}
	}
	if user.BookingID != "" {
		err = us.bookingDB.DeleteBooking(ctx, user.BookingID)
		if err != nil {
			return err
		}
	}
	return us.userDB.DeleteUser(ctx, id)
}
