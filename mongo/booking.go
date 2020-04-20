package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/yanrishbe/booking-service/model"
)

func (bs Booking) UpdateBooking(ctx context.Context, booking model.Booking) error {
	_id, err := primitive.ObjectIDFromHex(booking.ID)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", booking.ID, err)
	}
	query := bson.M{
		"_id": _id,
	}
	entity, err := booking.Entity()
	if err != nil {
		return err
	}
	updateDoc := bson.D{
		{"$set",
			bson.D{
				{"vip", entity.Vip},
				{"price", entity.Price},
				{"stars", entity.Persons},
				{"empty", entity.Empty},
				{"userId", entity.UserID},
				{"expiration", entity.Expiration},
				{"maxDays", entity.MaxDays},
			},
		},
	}
	_, err = bs.bookings.UpdateOne(ctx, query, updateDoc)
	if err != nil {
		return fmt.Errorf("could not update a booking %s", booking.ID)
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

func (bs Booking) GetAllBookings(ctx context.Context) ([]model.Booking, error) {
	cur, err := bs.bookings.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not find all bookings %v", err)
	}
	var bookingEntities []model.BookingEntity
	for cur.Next(ctx) {
		var bookingEntity model.BookingEntity
		err := cur.Decode(&bookingEntity)
		if err != nil {
			return nil, fmt.Errorf("could not decode mongo response %v", err)
		}
		bookingEntities = append(bookingEntities, bookingEntity)
	}
	defer func() {
		log.Println(cur.Close(ctx))
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
