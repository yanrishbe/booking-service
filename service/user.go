package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

type User struct {
	usersDB    UserRepository
	accountsDB AccountRepository
	bookingsDB BookingRepository
}

func NewUser(userDB UserRepository, accountDB AccountRepository, bookingDB BookingRepository) *User {
	return &User{
		usersDB:    userDB,
		accountsDB: accountDB,
		bookingsDB: bookingDB,
	}
}

func (us User) Create(ctx context.Context, user model.User) (string, error) {
	return us.usersDB.CreateUser(ctx, user)
}

func (us User) Login(ctx context.Context, loginRequest util.LoginRequest) (string, error) {
	passwordAndID, err := us.usersDB.GetPasswordAndID(ctx, loginRequest.Email)
	if err != nil {
		return "", err
	}
	if passwordAndID == nil {
		return "", fmt.Errorf("could not fetch users data")
	}
	return passwordAndID.ID, bcrypt.CompareHashAndPassword([]byte(passwordAndID.Password), []byte(loginRequest.Password))
}

func (us User) Get(ctx context.Context, id string) (*util.UserResponse, error) {
	user, err := us.usersDB.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	var account *model.Account
	if user.AccountID != "" {
		account, err = us.accountsDB.GetAccount(ctx, user.AccountID)
		if err != nil {
			return nil, err
		}
	}
	var booking *model.Booking
	if user.BookingID != "" {
		booking, err = us.bookingsDB.GetBooking(ctx, user.BookingID)
		if err != nil {
			return nil, err
		}
	}
	return util.NewUserResponse(user, account, booking)
}

func (us User) Update(ctx context.Context, userRequest util.UpdateUserRequest) error {
	return us.usersDB.UpdateUser(ctx, userRequest)
}

func (us User) Delete(ctx context.Context, id string) error {
	user, err := us.usersDB.GetUser(ctx, id)
	if err != nil {
		return err
	}
	if user.AccountID != "" {
		err = us.accountsDB.DeleteAccount(ctx, user.AccountID)
		if err != nil {
			return err
		}
	}
	if user.BookingID != "" {
		booking, err := us.bookingsDB.GetBooking(ctx, user.BookingID)
		if err != nil {
			return err
		}
		err = us.bookingsDB.UpdateBooking(ctx, model.Booking{
			ID:         booking.ID,
			Vip:        booking.Vip,
			Price:      booking.Price,
			Stars:      booking.Stars,
			Persons:    booking.Persons,
			Empty:      true,
			UserID:     nil,
			Expiration: nil,
			MaxDays:    0,
		})
		if err != nil {
			return err
		}
	}
	return us.usersDB.DeleteUser(ctx, id)
}

// supposed to be an inner call from bookings so unexported one
func (us User) DeleteAccount(ctx context.Context, accountID string, userID string) (int, error) {
	account, err := us.accountsDB.GetAccount(ctx, accountID)
	if err != nil {
		return 0, err
	}

	user, err := us.usersDB.GetUser(ctx, userID)
	if err != nil {
		return 0, err
	}

	err = us.accountsDB.DeleteAccount(ctx, user.AccountID)
	if err != nil {
		return 0, err
	}
	err = us.usersDB.UpdateAccountID(ctx, primitive.NilObjectID.Hex(), userID)
	if err != nil {
		return 0, err
	}
	return account.Amount, nil
}
