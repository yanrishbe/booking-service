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

// todo add booking for a date not from today
func (bk Booking) Create(ctx context.Context, newBookingReq util.BookingRequest, userID string) error {
	user, err := bk.usersDB.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	account, err := bk.accountsDB.GetAccount(ctx, user.AccountID)
	if err != nil {
		return err
	}

	oldBooking, err := bk.bookingsDB.GetBooking(ctx, newBookingReq.ID)
	if err != nil {
		return err
	}

	price := oldBooking.Price * newBookingReq.MaxDays
	if account.Amount < price {
		return fmt.Errorf("insufficient funds to book an apartment")
	}

	if !oldBooking.Empty {
		return fmt.Errorf("the apartment is already booked")
	}

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
	err = bk.accountsDB.UpdateAccount(ctx, *adminAccount)
	if err != nil {
		return err
	}

	err = bk.usersDB.UpdateBookingID(ctx, newBookingReq.ID, userID)
	if err != nil {
		return err
	}

	expirationTime := time.Now().AddDate(0, 0, newBookingReq.MaxDays)
	return bk.bookingsDB.UpdateBooking(ctx, model.Booking{
		ID:         oldBooking.ID,
		Vip:        oldBooking.Vip,
		Price:      oldBooking.Price,
		Stars:      oldBooking.Stars,
		Persons:    oldBooking.Persons,
		Empty:      false,
		UserID:     &userID,
		Expiration: &expirationTime,
		MaxDays:    newBookingReq.MaxDays,
	})
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

// todo return refund for the expired bookings
func (bk Booking) Delete(ctx context.Context, bookingID string, userID string) error {
	oldBooking, err := bk.bookingsDB.GetBooking(ctx, bookingID)
	if err != nil {
		return err
	}
	err = bk.bookingsDB.UpdateBooking(ctx, model.Booking{
		ID:         oldBooking.ID,
		Vip:        oldBooking.Vip,
		Price:      oldBooking.Price,
		Stars:      oldBooking.Stars,
		Persons:    oldBooking.Persons,
		Empty:      true,
		UserID:     nil,
		Expiration: nil,
		MaxDays:    0,
	})
	if err != nil {
		return err
	}

	err = bk.usersDB.UpdateBookingID(ctx, "", userID)
	if err != nil {
		return err
	}

	return nil
}

func (bk Booking) Update(ctx context.Context, newBookingReq util.BookingRequest, bookingID string, userID string) error {
	user, err := bk.usersDB.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	account, err := bk.accountsDB.GetAccount(ctx, user.AccountID)
	if err != nil {
		return err
	}

	oldBooking, err := bk.bookingsDB.GetBooking(ctx, bookingID)
	if err != nil {
		return err
	}

	price := oldBooking.Price * newBookingReq.MaxDays
	if account.Amount < price {
		return fmt.Errorf("insufficient funds to update the booking")
	}

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
	err = bk.accountsDB.UpdateAccount(ctx, *adminAccount)
	if err != nil {
		return err
	}

	err = bk.usersDB.UpdateBookingID(ctx, newBookingReq.ID, userID)
	if err != nil {
		return err
	}

	expirationTime := oldBooking.Expiration.AddDate(0, 0, newBookingReq.MaxDays)
	return bk.bookingsDB.UpdateBooking(ctx, model.Booking{
		ID:         oldBooking.ID,
		Vip:        oldBooking.Vip,
		Price:      oldBooking.Price,
		Stars:      oldBooking.Stars,
		Persons:    oldBooking.Persons,
		Empty:      false,
		UserID:     &userID,
		Expiration: &expirationTime,
		MaxDays:    newBookingReq.MaxDays,
	})
}
