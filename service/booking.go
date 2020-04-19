package service

import (
	"context"
	"fmt"
	"time"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

type Booking struct {
	bookingsDB BookingRepository
	accountsDB AccountRepository
	usersDB    UserRepository
}

func NewBooking(bookingsDB BookingRepository, accountsDB AccountRepository, usersDB UserRepository) *Booking {
	return &Booking{
		bookingsDB: bookingsDB,
		accountsDB: accountsDB,
		usersDB:    usersDB,
	}
}

func (bk Booking) Create(ctx context.Context, newBooking model.Booking, userID string) error {
	user, err := bk.usersDB.GetUser(ctx, userID)
	if err != nil {
		return err
	}
	account, err := bk.accountsDB.GetAccount(ctx, user.AccountID)
	if err != nil {
		return err
	}
	price := newBooking.Price * newBooking.MaxDays
	if account.Amount < price {
		return fmt.Errorf("insufficient funds to book an apartment")
	}
	oldBooking, err := bk.bookingsDB.GetBooking(ctx, newBooking.ID)
	if err != nil {
		return err
	}
	if !oldBooking.Empty {
		return fmt.Errorf("the apartment is already booked")
	}
	newBooking.Empty = false
	expirationTime := time.Now().AddDate(0, 0, newBooking.MaxDays)
	newBooking.Expiration = &expirationTime

	account.Amount -= price
	err = bk.accountsDB.UpdateAccount(ctx, *account)
	if err != nil {
		return err
	}

	adminAccount, err := bk.accountsDB.GetAdminAccount(ctx)
	if err != nil {
		return err
	}
	adminAccount.Amount += price
	err = bk.accountsDB.UpdateAccount(ctx, *account)
	if err != nil {
		return err
	}

	return bk.bookingsDB.UpdateBooking(ctx, newBooking)
}

func (bk Booking) Get(ctx context.Context, id string) (*model.Booking, error) {
	return bk.bookingsDB.GetBooking(ctx, id)
}

func (bk Booking) GetAll(ctx context.Context) ([]util.GetAllBookingsResponse, error) {
	bookings, err := bk.bookingsDB.GetAllBookings(ctx)
	if err != nil {
		return nil, err
	}
	resp := util.NewGetAllBookingsResponse(bookings)
	return resp, nil
}

// todo update and delete
// func (ac Account) Update(ctx context.Context, newAccount model.Account, accountID string, userID string) error {
// 	oldAccount, err := ac.accountsDB.GetAccount(ctx, accountID)
// 	if err != nil {
// 		return err
// 	}
// 	amount := oldAccount.Amount + newAccount.Amount
// 	switch {
// 	case oldAccount.UserID != userID:
// 		{
// 			return fmt.Errorf("could not update account: the user does does not have enough rights")
// 		}
// 	case amount < 0:
// 		{
// 			oldAccount.BlockedCounter++
// 			if oldAccount.BlockedCounter == 10 {
// 				oldAccount.Blocked = true
// 			}
// 			defer func() {
// 				log.Println(ac.accountsDB.UpdateAccount(ctx, *oldAccount))
// 			}()
// 			return fmt.Errorf("could not update account: new amount is less then the current one")
// 		}
// 	}
// 	newAccount.Amount = amount
// 	return ac.accountsDB.UpdateAccount(ctx, newAccount)
// }
