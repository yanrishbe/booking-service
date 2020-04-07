package mongo

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/yanrishbe/booking-service/model"
)

type Booking struct {
	*mongo.Database
	users    *mongo.Collection
	bookings *mongo.Collection
	accounts *mongo.Collection
}

func NewBooking(ctx context.Context) (*Booking, error) {
	connStr, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		return nil, errors.New("empty connection string")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for retries := 5; retries >= 0; retries-- {
		err = client.Ping(context.Background(), readpref.Primary())
		if err == nil {
			break
		}
		log.Printf("reconnecting to db: %v", err)
		time.Sleep(time.Second * 2)
	}

	bookingServiceDB := client.Database("booking-service")
	users := bookingServiceDB.Collection("users")
	bookings := bookingServiceDB.Collection("bookings")
	accounts := bookingServiceDB.Collection("accounts")

	uIndexes := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{
				Key:   "bookingId",
				Value: bsonx.Int64(1),
			}},
			Options: options.Index().SetName("bookingId"),
		},
		{
			Keys: bsonx.Doc{{
				Key:   "accountId",
				Value: bsonx.Int64(1),
			}},
			Options: options.Index().SetName("accountId"),
		},
	}
	_, err = users.Indexes().CreateMany(ctx, uIndexes)
	if err != nil {
		return nil, err
	}
	return &Booking{
		Database: bookingServiceDB,
		users:    users,
		bookings: bookings,
		accounts: accounts,
	}, err
}

func (bs Booking) Create(ctx context.Context, user model.User) (string, error) {
	query := bson.M{
		"email": primitive.Regex{Pattern: user.Email, Options: "i"},
	}
	count, err := bs.users.CountDocuments(ctx, query)
	if err != nil {
		return "", errors.New("count error " + err.Error())
	}
	if count > 0 {
		return "", errors.New("login already exists: " + user.Email)
	}
	res, err := bs.users.InsertOne(ctx, user)
	if err != nil {
		return "", errors.New("couldn't insert a user")
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
