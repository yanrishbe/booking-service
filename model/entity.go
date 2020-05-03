package model

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const Admin = "yana.strabuk@gmail.com"
const AdminBank = "JPMorgan"

type User struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	AccountID  string `json:"accountId,omitempty" bson:"accountId"`
	BookingID  string `json:"bookingId,omitempty" bson:"bookingId"`
	Name       string `json:"name" bson:"name"`
	Surname    string `json:"surname" bson:"surname"`
	Patronymic string `json:"patronymic" bson:"patronymic"`
	Phone      string `json:"phone" bson:"phone" `
	Email      string `json:"email" bson:"email"`
	Password   string `json:"password" bson:"password"`
}

func (u User) Entity() (*UserEntity, error) {
	entity := UserEntity{
		Name:       u.Name,
		Surname:    u.Surname,
		Patronymic: u.Patronymic,
		Phone:      u.Phone,
		Email:      u.Email,
		Password:   u.Password,
	}
	var err error
	if u.ID != "" {
		entity.ID, err = primitive.ObjectIDFromHex(u.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
	}
	if u.AccountID != "" {
		objID, err := primitive.ObjectIDFromHex(u.AccountID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
		entity.AccountID = &objID
	}
	if u.BookingID != "" {
		objID, err := primitive.ObjectIDFromHex(u.BookingID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
		entity.BookingID = &objID
	}
	return &entity, nil
}

type UserEntity struct {
	ID         primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	AccountID  *primitive.ObjectID `json:"accountId,omitempty" bson:"accountId"`
	BookingID  *primitive.ObjectID `json:"bookingId,omitempty" bson:"bookingId"`
	Name       string              `json:"name" bson:"name"`
	Surname    string              `json:"surname" bson:"surname"`
	Patronymic string              `json:"patronymic" bson:"patronymic"`
	Phone      string              `json:"phone" bson:"phone" `
	Email      string              `json:"email" bson:"email"`
	Password   string              `json:"password" bson:"password"`
}

func (ue UserEntity) DTO() User {
	user := User{
		ID:         ue.ID.Hex(),
		Name:       ue.Name,
		Surname:    ue.Surname,
		Patronymic: ue.Patronymic,
		Phone:      ue.Phone,
		Email:      ue.Email,
		Password:   ue.Password,
	}
	if ue.AccountID != nil {
		user.AccountID = ue.AccountID.Hex()
	}
	if ue.BookingID != nil {
		user.BookingID = ue.BookingID.Hex()
	}
	return user
}

type Account struct {
	ID             string `json:"id,omitempty" bson:"_id,omitempty"`
	UserID         string `json:"userId" bson:"userId"`
	CreditCard     bool   `json:"creditCard" bson:"creditCard"`
	LegalEntity    bool   `json:"legalEntity" bson:"legalEntity"`
	Blocked        bool   `json:"blocked" bson:"blocked"`
	BlockedCounter int    `bson:"blockedCounter"`
	Bank           string `json:"bank" bson:"bank"`
	Amount         int    `json:"amount" bson:"amount"`
}

func (a Account) Entity() (*AccountEntity, error) {
	entity := AccountEntity{
		CreditCard:     a.CreditCard,
		LegalEntity:    a.LegalEntity,
		Blocked:        a.Blocked,
		BlockedCounter: a.BlockedCounter,
		Bank:           a.Bank,
		Amount:         a.Amount,
	}
	var err error
	if a.ID != "" {
		entity.ID, err = primitive.ObjectIDFromHex(a.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
	}
	if a.UserID != "" {
		entity.UserID, err = primitive.ObjectIDFromHex(a.UserID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
	}
	return &entity, nil
}

type AccountEntity struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID         primitive.ObjectID `json:"userId" bson:"userId"`
	CreditCard     bool               `json:"creditCard" bson:"creditCard"`
	LegalEntity    bool               `json:"legalEntity" bson:"legalEntity"`
	Blocked        bool               `json:"blocked" bson:"blocked"`
	BlockedCounter int                `bson:"blockedCounter"`
	Bank           string             `json:"bank" bson:"bank"`
	Amount         int                `json:"amount" bson:"amount"`
}

func (ae AccountEntity) DTO() Account {
	return Account{
		ID:             ae.ID.Hex(),
		UserID:         ae.UserID.Hex(),
		CreditCard:     ae.CreditCard,
		LegalEntity:    ae.LegalEntity,
		Blocked:        ae.Blocked,
		BlockedCounter: ae.BlockedCounter,
		Bank:           ae.Bank,
		Amount:         ae.Amount,
	}
}

type Booking struct {
	ID         string     `json:"id,omitempty" bson:"_id,omitempty"`
	Vip        bool       `json:"vip" bson:"vip"`
	Price      int        `json:"price" bson:"price"`
	Stars      int        `json:"stars" bson:"stars"`
	Persons    int        `json:"persons" bson:"persons"`
	Empty      bool       `json:"empty" bson:"empty"`
	UserID     *string    `json:"userId" bson:"userId"`
	Expiration *time.Time `json:"expiration" bson:"expiration"`
	MaxDays    int        `json:"maxDays" bson:"maxDays"`
	URL        string     `json:"url" bson:"url"`
}

func (b Booking) Entity() (*BookingEntity, error) {
	entity := BookingEntity{
		Vip:        b.Vip,
		Price:      b.Price,
		Stars:      b.Stars,
		Persons:    b.Persons,
		Empty:      b.Empty,
		Expiration: b.Expiration,
		MaxDays:    b.MaxDays,
		URL:        b.URL,
	}
	var err error
	if b.ID != "" {
		entity.ID, err = primitive.ObjectIDFromHex(b.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
	}
	if b.UserID != nil {
		objID, err := primitive.ObjectIDFromHex(*b.UserID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
		entity.UserID = &objID
	}
	return &entity, nil
}

type BookingEntity struct {
	ID         primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Vip        bool                `json:"vip" bson:"vip"`
	Price      int                 `json:"price" bson:"price"`
	Stars      int                 `json:"stars" bson:"stars"`
	Persons    int                 `json:"persons" bson:"persons"`
	Empty      bool                `json:"empty" bson:"empty"`
	UserID     *primitive.ObjectID `json:"userId" bson:"userId"`
	Expiration *time.Time          `json:"expiration" bson:"expiration"`
	MaxDays    int                 `json:"maxDays" bson:"maxDays"`
	URL        string              `json:"url" bson:"url"`
}

func (be BookingEntity) DTO() Booking {
	booking := Booking{
		ID:         be.ID.Hex(),
		Vip:        be.Vip,
		Price:      be.Price,
		Stars:      be.Stars,
		Persons:    be.Persons,
		Empty:      be.Empty,
		Expiration: be.Expiration,
		MaxDays:    be.MaxDays,
		URL:        be.URL,
	}
	if be.UserID != nil {
		userID := be.UserID.Hex()
		booking.UserID = &userID
	}
	return booking
}
