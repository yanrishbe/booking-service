package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yanrishbe/booking-service/model"
)

func (bs Booking) CreateBooking(ctx context.Context, booking model.Booking) (string, error) {
	_id, err := primitive.ObjectIDFromHex(booking.ID)
	if err != nil {
		return "", fmt.Errorf("could not parse object id %s: %v", booking.ID, err)
	}
	query := bson.M{
		"_id": _id,
	}
	count, err := bs.bookings.CountDocuments(ctx, query)
	if err != nil {
		return "", fmt.Errorf("count error %v", err)
	}
	if count > 0 {
		return "", fmt.Errorf("booking already exists %s", booking.ID)
	}
	res, err := bs.bookings.InsertOne(ctx, booking)
	if err != nil {
		return "", fmt.Errorf("couldn't create a booking %s: %v", booking.ID, err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (bs Booking) UpdateBooking(ctx context.Context, booking model.Booking) error {
	_id, err := primitive.ObjectIDFromHex(booking.ID)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", booking.ID, err)
	}
	query := bson.M{
		"_id": _id,
	}
	bookingEntity, err := booking.Entity()
	if err != nil {
		return err
	}
	_, err = bs.bookings.UpdateOne(ctx, query, bookingEntity)
	if err != nil {
		return fmt.Errorf("could not update a booking %s", booking.ID)
	}
	return nil
}

func (bs Booking) DeleteBooking(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", id, err)
	}
	query := bson.M{
		"_id": _id,
	}
	_, err = bs.bookings.DeleteOne(ctx, query)
	if err != nil {
		return fmt.Errorf("could not delete a booking %s", id)
	}
	return nil
}

func (bs Booking) GetBooking(ctx context.Context, id string) (*model.Booking, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("could not parse object id %s: %v", id, err)
	}
	query := bson.M{
		"_id": _id,
	}
	response := bs.bookings.FindOne(ctx, query)
	if response.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	var bookingEntity model.BookingEntity
	err = response.Decode(&bookingEntity)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongo response %v", err)
	}
	booking := bookingEntity.DTO()
	return &booking, nil
}
