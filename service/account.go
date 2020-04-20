package service

import (
	"context"
	"fmt"
	"log"

	"github.com/yanrishbe/booking-service/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (ac Account) Create(ctx context.Context, accountRequest util.AccountRequest, userID string) (string, error) {
	account, err := util.AccountFromRequest(accountRequest)
	if err != nil {
		return "", err
	}

	if account.Amount < 0 {
		return "", fmt.Errorf("insufficient funds to create an account")
	}
	user, err := ac.usersDB.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("the user does not exist")
	}
	if user.AccountID != "" && user.AccountID != primitive.NilObjectID.Hex() {
		return "", fmt.Errorf("could not create an account: an account for the user exists")
	}

	account.UserID = userID
	accountID, err := ac.accountsDB.CreateAccount(ctx, *account)
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

func (ac Account) Get(ctx context.Context, id string) (*util.AccountResponse, error) {
	account, err := ac.accountsDB.GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	accResponse := util.NewAccountResponse(*account)
	return accResponse, nil
}

func (ac Account) Update(ctx context.Context, newAccountReq util.AccountRequest, accountID string, userID string) error {
	newAccount, err := util.AccountFromRequest(newAccountReq)
	if err != nil {
		return err
	}

	oldAccount, err := ac.accountsDB.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}
	amount := oldAccount.Amount + newAccount.Amount
	switch {
	case oldAccount.UserID != userID:
		{
			return fmt.Errorf("could not update account: the user does does not have enough rights")
		}
	case amount < 0:
		{
			oldAccount.BlockedCounter++
			if oldAccount.BlockedCounter == 10 {
				oldAccount.Blocked = true
			}
			defer func() {
				log.Println(ac.accountsDB.UpdateAccount(ctx, *oldAccount))
			}()
			return fmt.Errorf("could not update account: new amount is less then the current one")
		}
	}
	newAccount.Amount = amount
	newAccount.ID = oldAccount.ID
	return ac.accountsDB.UpdateAccount(ctx, *newAccount)
}
