package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const Admin = "yana.strabuk@gmail.com"
const AdminBank = "JPMorgan"

type User struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	AccountID  string `json:"accountId,omitempty" bson:"accountId,omitempty"`
	BookingID  string `json:"bookingId,omitempty" bson:"bookingId,omitempty"`
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
	AccountID  *primitive.ObjectID `json:"accountId,omitempty" bson:"accountId,omitempty"`
	BookingID  *primitive.ObjectID `json:"bookingId,omitempty" bson:"bookingId,omitempty"`
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
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Bank   string `json:"bank" bson:"bank"`
	Amount int    `json:"amount" bson:"amount"`
}

func (a Account) Entity() (*AccountEntity, error) {
	entity := AccountEntity{
		Bank:   a.Bank,
		Amount: a.Amount,
	}
	var err error
	if a.ID != "" {
		entity.ID, err = primitive.ObjectIDFromHex(a.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
	}
	return &entity, nil
}

type AccountEntity struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Bank   string             `json:"bank" bson:"bank"`
	Amount int                `json:"amount" bson:"amount"`
}

func (ae AccountEntity) DTO() Account {
	return Account{
		ID:     ae.ID.Hex(),
		Bank:   ae.Bank,
		Amount: ae.Amount,
	}
}

type Booking struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Vip     bool   `json:"vip" bson:"vip"`
	Price   int    `json:"price" bson:"price"`
	Stars   int    `json:"stars" bson:"stars"`
	Persons int    `json:"persons" bson:"persons"`
}

func (b Booking) Entity() (*BookingEntity, error) {
	entity := BookingEntity{
		Vip:     b.Vip,
		Price:   b.Price,
		Stars:   b.Stars,
		Persons: b.Persons,
	}
	var err error
	if b.ID != "" {
		entity.ID, err = primitive.ObjectIDFromHex(b.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid id object id: %v", err)
		}
	}
	return &entity, nil
}

type BookingEntity struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Vip     bool               `json:"vip" bson:"vip"`
	Price   int                `json:"price" bson:"price"`
	Stars   int                `json:"stars" bson:"stars"`
	Persons int                `json:"persons" bson:"persons"`
}

func (be BookingEntity) DTO() Booking {
	return Booking{
		ID:      be.ID.Hex(),
		Vip:     be.Vip,
		Price:   be.Price,
		Stars:   be.Stars,
		Persons: be.Persons,
	}
}
