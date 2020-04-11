package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

type User struct {
	userDB    UserRepository
	accountDB AccountRepository
	bookingDB BookingRepository
}

func NewUser(repository UserRepository) *User {
	return &User{userDB: repository}
}

func (us User) Create(ctx context.Context, user model.User) (string, error) {
	return us.userDB.CreateUser(ctx, user)
}

func (us User) Login(ctx context.Context, loginRequest util.LoginRequest) error {
	password, err := us.userDB.CheckPassword(ctx, loginRequest.Email)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(loginRequest.Password))
}

func (us User) Get(ctx context.Context, email string) (*util.UserResponse, error) {
	user, err := us.userDB.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	account, err := us.accountDB.GetAccount(ctx, user.AccountID)
	if err != nil {
		return nil, err
	}
	booking, err := us.bookingDB.GetBooking(ctx, user.BookingID)
	if err != nil {
		return nil, err
	}
	return util.NewUserResponse(user, account, booking)
}
