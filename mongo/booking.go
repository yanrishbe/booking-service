package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/yanrishbe/booking-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var bookingEntity model.BookingEntity
	err = response.Decode(&bookingEntity)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongo response %v", err)
	}
	booking := bookingEntity.DTO()
	return &booking, nil
}

func (bs Booking) GetAllBookings(ctx context.Context) ([]model.Booking, error) {
	cur, err := bs.bookings.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not find all bookings %v", err)
	}
	var bookingEntities []model.BookingEntity
	for cur.Next(context.TODO()) {
		var bookingEntity model.BookingEntity
		err := cur.Decode(&bookingEntity)
		if err != nil {
			return nil, fmt.Errorf("could not decode mongo response %v", err)
		}
		bookingEntities = append(bookingEntities, bookingEntity)
	}
	defer func() {
		log.Fatalln(cur.Close(ctx))
	}()
	err = cur.Err()
	if err != nil {
		return nil, fmt.Errorf("cursor error %v", err)
	}
	var bookings []model.Booking
	for i := range bookingEntities {
		bookings = append(bookings, bookingEntities[i].DTO())
	}
	return bookings, nil
}
