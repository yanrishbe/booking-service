package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

type Account struct {
	accountDB AccountRepository
	usersDB   UserRepository
}

func NewAccount(accountDB AccountRepository, usersDB UserRepository) *Account {
	return &Account{
		accountDB: accountDB,
		usersDB:   usersDB,
	}
}

func (ac Account) Create(ctx context.Context, account model.Account, userID string) (string, error) {
	if account.Amount < 0 {
		return "", fmt.Errorf("insufficient funds to create an account")
	}
	user, err := ac.usersDB.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user.AccountID != "" && user.AccountID != primitive.NilObjectID.Hex() {
		return "", fmt.Errorf("could not create an account: an account for the user exists")
	}
	accountID, err := ac.accountDB.CreateAccount(ctx, account)
	if err != nil {
		return "", err
	}
	user.AccountID = accountID
	err = ac.usersDB.UpdateAccountID(ctx, accountID, user.ID)
	if err != nil {
		return "", err
	}
	return accountID, nil
}

// func (ac Account) Get(ctx context.Context, id string) (*util.UserResponse, error) {
// 	user, err := us.userDB.GetUser(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var account *model.Account
// 	if user.AccountID != "" {
// 		account, err = us.accountDB.GetAccount(ctx, user.AccountID)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	var booking *model.Booking
// 	if user.BookingID != "" {
// 		booking, err = us.bookingDB.GetBooking(ctx, user.BookingID)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return util.NewUserResponse(user, account, booking)
// }
//
// func (ac Account) Update(ctx context.Context, userRequest util.UpdateUserRequest) error {
// 	return us.userDB.UpdateUser(ctx, userRequest)
// }
//
// func (ac Account) Delete(ctx context.Context, id string) error {
// 	user, err := us.userDB.GetUser(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	if user.AccountID != "" {
// 		err = us.accountDB.DeleteAccount(ctx, user.AccountID)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	if user.BookingID != "" {
// 		err = us.bookingDB.DeleteBooking(ctx, user.BookingID)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return us.userDB.DeleteUser(ctx, id)
// }
