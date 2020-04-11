package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"golang.org/x/crypto/bcrypt"

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

func (bs Booking) CreateUser(ctx context.Context, user model.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		return "", fmt.Errorf("could not hash the password")
	}
	user.Password = string(hash)
	query := bson.M{
		"email": user.Email,
	}
	count, err := bs.users.CountDocuments(ctx, query)
	if err != nil {
		return "", fmt.Errorf("count error %v", err)
	}
	if count > 0 {
		return "", fmt.Errorf("login already exists %s", user.Email)
	}
	res, err := bs.users.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("couldn't create a user %s: %v", user.Email, err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

type password struct {
	Password string `bson:"password"`
}

func (bs Booking) CheckPassword(ctx context.Context, email string) (string, error) {
	opts := options.FindOne().SetProjection(bson.M{
		"password": 1,
	})
	query := bson.M{
		"email": email,
	}
	var passw password
	err := bs.users.FindOne(ctx, query, opts).Decode(&passw)
	if err != nil {
		return "", fmt.Errorf("could not decode mongo response %v", err)
	}
	return passw.Password, nil
}

func (bs Booking) UpdateUser(ctx context.Context, user model.User) error {
	query := bson.M{
		"email": user.Email,
	}
	userEntity, err := user.Entity()
	if err != nil {
		return err
	}
	_, err = bs.users.UpdateOne(ctx, query, userEntity)
	if err != nil {
		return fmt.Errorf("could not update a user %s", user.Email)
	}
	return nil
}

func (bs Booking) DeleteUser(ctx context.Context, email string) error {
	query := bson.M{
		"email": email,
	}
	_, err := bs.users.DeleteOne(ctx, query)
	if err != nil {
		return fmt.Errorf("could not delete a user %s", email)
	}
	return nil
}

func (bs Booking) GetUser(ctx context.Context, email string) (*model.User, error) {
	query := bson.M{
		"email": email,
	}
	response := bs.users.FindOne(ctx, query)
	var userEntity model.UserEntity
	err := response.Decode(&userEntity)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongo response %v", err)
	}
	user := userEntity.DTO()
	return &user, nil
}

func (bs Booking) GetAllUsers(ctx context.Context) ([]model.User, error) {
	cur, err := bs.users.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not find all users %v", err)
	}
	var userEntities []model.UserEntity
	for cur.Next(context.TODO()) {
		var userEntity model.UserEntity
		err := cur.Decode(&userEntity)
		if err != nil {
			return nil, fmt.Errorf("could not decode mongo response %v", err)
		}
		userEntities = append(userEntities, userEntity)
	}
	defer func() {
		log.Fatalln(cur.Close(ctx))
	}()
	err = cur.Err()
	if err != nil {
		return nil, fmt.Errorf("cursor error %v", err)
	}
	var users []model.User
	for i := range userEntities {
		users = append(users, userEntities[i].DTO())
	}
	return users, nil
}
