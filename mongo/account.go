package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yanrishbe/booking-service/model"
)

func (bs Booking) CreateAccount(ctx context.Context, account model.Account) (string, error) {
	entity, err := account.Entity()
	if err != nil {
		return "", err
	}
	res, err := bs.accounts.InsertOne(ctx, entity)
	if err != nil {
		return "", fmt.Errorf("couldn't create an account %s: %v", account.ID, err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (bs Booking) UpdateAccount(ctx context.Context, account model.Account) error {
	_id, err := primitive.ObjectIDFromHex(account.ID)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", account.ID, err)
	}
	query := bson.M{
		"_id": _id,
	}
	updateDoc := bson.D{
		{"$set",
			bson.D{
				{"bank", account.Bank},
				{"amount", account.Amount},
				{"creditCard", account.CreditCard},
				{"legalEntity", account.LegalEntity},
				{"blocked", account.Blocked},
				{"blockedCounter", account.BlockedCounter},
			},
		},
	}
	_, err = bs.accounts.UpdateOne(ctx, query, updateDoc)
	if err != nil {
		return fmt.Errorf("could not update an account %s", account.ID)
	}
	return nil
}

func (bs Booking) GetAccount(ctx context.Context, id string) (*model.Account, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("could not parse object id %s: %v", id, err)
	}
	query := bson.M{
		"_id": _id,
	}
	response := bs.accounts.FindOne(ctx, query)
	if response.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	var accountEntity model.AccountEntity
	err = response.Decode(&accountEntity)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongo response %v", err)
	}
	account := accountEntity.DTO()
	return &account, nil
}

func (bs Booking) GetAdminAccount(ctx context.Context) (*model.Account, error) {
	query := bson.M{
		"email": model.Admin,
	}
	response := bs.users.FindOne(ctx, query)
	if response.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	var userEntity model.UserEntity
	err := response.Decode(&userEntity)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongo response %v", err)
	}

	query = bson.M{
		"_id": userEntity.AccountID,
	}
	response = bs.accounts.FindOne(ctx, query)
	if response.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	var accountEntity model.AccountEntity
	err = response.Decode(&accountEntity)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongo response %v", err)
	}
	account := accountEntity.DTO()
	return &account, nil
}

func (bs Booking) DeleteAccount(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", id, err)
	}
	query := bson.M{
		"_id": _id,
	}
	_, err = bs.accounts.DeleteOne(ctx, query)
	if err != nil {
		return fmt.Errorf("could not delete an account %s", id)
	}
	return nil
}
