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

	account.UserID = userID
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

func (ac Account) Update(ctx context.Context, newAccount model.Account, accountID string, userID string) error {
	oldAccount, err := ac.accountsDB.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	switch {
	case oldAccount.UserID != userID:
		{
			return fmt.Errorf("could not update account: the user does does not have enough rights")
		}
	case newAccount.Amount < oldAccount.Amount:
		{
			oldAccount.BlockedCounter++
			if oldAccount.BlockedCounter == 10 {
				oldAccount.Blocked = true
			}
			defer func() error {
				return ac.accountsDB.UpdateAccount(ctx, *oldAccount)
			}()
			return fmt.Errorf("could not update account: new amount is less then the current one")
		}
	}

	return ac.accountsDB.UpdateAccount(ctx, newAccount)
}
