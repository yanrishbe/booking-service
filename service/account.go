package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yanrishbe/booking-service/model"
)

type Account struct {
	accountsDB AccountRepository
	usersDB    UserRepository
}

func NewAccount(accountDB AccountRepository, usersDB UserRepository) *Account {
	return &Account{
		accountsDB: accountDB,
		usersDB:    usersDB,
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
	accountID, err := ac.accountsDB.CreateAccount(ctx, account)
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

func (ac Account) Get(ctx context.Context, id string) (*model.Account, error) {
	return ac.accountsDB.GetAccount(ctx, id)
}

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
// 		err = us.accountsDB.DeleteAccount(ctx, user.AccountID)
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
