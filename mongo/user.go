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
	"github.com/yanrishbe/booking-service/util"
)

type Booking struct {
	*mongo.Database
	users    *mongo.Collection
	bookings *mongo.Collection
	accounts *mongo.Collection
}

const adminPassword = "admin"

// todo add transactions may be one day
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

	err = initDatabase(ctx, users, bookings, accounts)
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

func initDatabase(ctx context.Context, users *mongo.Collection, bookings *mongo.Collection, accounts *mongo.Collection) error {
	userId := primitive.NewObjectID()
	res, err := accounts.InsertOne(ctx, bson.D{
		{"bank", model.AdminBank},
		{"amount", 5000000},
		{"userId", userId},
	})
	if err != nil {
		return fmt.Errorf("couldn't create admin account%v", err)
	}
	accountID := res.InsertedID.(primitive.ObjectID)
	hash, err := bcrypt.GenerateFromPassword([]byte(adminPassword), 5)
	if err != nil {
		return fmt.Errorf("could not hash the admin password")
	}
	_, err = users.InsertOne(ctx, bson.D{
		{"_id", userId},
		{"email", model.Admin},
		{"password", string(hash)},
		{"accountId", accountID},
	})

	if err != nil {
		return fmt.Errorf("couldn't create admin %v", err)
	}

	rooms := []model.BookingEntity{
		{
			Vip:     true,
			Empty:   true,
			Price:   10000,
			Stars:   5,
			Persons: 3,
		},
		{
			Vip:     false,
			Empty:   true,
			Price:   7000,
			Stars:   3,
			Persons: 2,
		},
		{
			Vip:     false,
			Empty:   true,
			Price:   8000,
			Stars:   3,
			Persons: 3,
		},
		{
			Vip:     false,
			Empty:   true,
			Price:   5000,
			Stars:   1,
			Persons: 1,
		},
		{
			Vip:     true,
			Empty:   true,
			Price:   20000,
			Stars:   5,
			Persons: 2,
		},
		{
			Vip:     true,
			Empty:   true,
			Price:   12000,
			Stars:   4,
			Persons: 2,
		},
		{
			Vip:     false,
			Empty:   true,
			Price:   6500,
			Stars:   3,
			Persons: 3,
		},
		{
			Vip:     false,
			Empty:   true,
			Price:   7500,
			Stars:   3,
			Persons: 1,
		},
		{
			Vip:     false,
			Empty:   true,
			Price:   10000,
			Stars:   4,
			Persons: 4,
		},
		{
			Vip:     true,
			Empty:   true,
			Price:   19000,
			Stars:   5,
			Persons: 2,
		},
	}
	interfaceList := make([]interface{}, len(rooms))
	for i := range rooms {
		interfaceList[i] = rooms[i]
	}
	_, err = bookings.InsertMany(ctx, interfaceList)
	if err != nil {
		return fmt.Errorf("couldn't create a booking %v", err)
	}
	return nil
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
	entity, err := user.Entity()
	if err != nil {
		return "", err
	}
	res, err := bs.users.InsertOne(ctx, entity)
	if err != nil {
		return "", fmt.Errorf("couldn't create a user %s: %v", user.Email, err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

type passwordAndID struct {
	Password string              `bson:"password"`
	ID       *primitive.ObjectID `bson:"_id"`
}

type PasswordAndIDResponse struct {
	Password string
	ID       string
}

func (bs Booking) GetPasswordAndID(ctx context.Context, email string) (*PasswordAndIDResponse, error) {
	opts := options.FindOne().SetProjection(bson.M{
		"password": 1,
		"_id":      1,
	})
	query := bson.M{
		"email": email,
	}
	var passw passwordAndID
	response := bs.users.FindOne(ctx, query, opts)
	if response.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	err := response.Decode(&passw)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongo response %v", err)
	}
	passwordAndIDResponse := PasswordAndIDResponse{
		Password: passw.Password,
	}
	if passw.ID != nil {
		passwordAndIDResponse.ID = passw.ID.Hex()
	}
	return &passwordAndIDResponse, nil
}

func (bs Booking) GetUser(ctx context.Context, id string) (*model.User, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("could not parse object id %s: %v", id, err)
	}
	query := bson.M{
		"_id": _id,
	}
	response := bs.users.FindOne(ctx, query)
	if response.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	var userEntity model.UserEntity
	err = response.Decode(&userEntity)
	if err != nil {
		return nil, fmt.Errorf("could not decode mongo response %v", err)
	}
	user := userEntity.DTO()
	return &user, nil
}

func (bs Booking) UpdateUser(ctx context.Context, userRequest util.UpdateUserRequest) error {
	_id, err := primitive.ObjectIDFromHex(userRequest.ID)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", userRequest.ID, err)
	}
	query := bson.M{
		"_id": _id,
	}
	updateDoc := bson.D{
		{"$set",
			bson.D{
				{"name", userRequest.Name},
				{"surname", userRequest.Surname},
				{"patronymic", userRequest.Patronymic},
				{"phone", userRequest.Phone},
				{"email", userRequest.Email},
			},
		},
	}
	_, err = bs.users.UpdateOne(ctx, query, updateDoc)
	if err != nil {
		return fmt.Errorf("could not update a user %s", userRequest.ID)
	}
	return nil
}

func (bs Booking) UpdateAccountID(ctx context.Context, accID string, userID string) error {
	_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", userID, err)
	}

	var _accID *primitive.ObjectID
	if accID != "" {
		val, err := primitive.ObjectIDFromHex(accID)
		if err != nil {
			return fmt.Errorf("could not parse object id %s: %v", accID, err)
		}
		_accID = &val
	}

	query := bson.M{
		"_id": _id,
	}
	updateDoc := bson.D{
		{"$set",
			bson.D{
				{"accountId", _accID},
			},
		},
	}
	_, err = bs.users.UpdateOne(ctx, query, updateDoc)
	if err != nil {
		return fmt.Errorf("could not update user's accountId %s", userID)
	}
	return nil
}

func (bs Booking) UpdateBookingID(ctx context.Context, bookID string, userID string) error {
	_id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", userID, err)
	}
	var _bookID *primitive.ObjectID
	if bookID != "" {
		val, err := primitive.ObjectIDFromHex(bookID)
		if err != nil {
			return fmt.Errorf("could not parse object id %s: %v", bookID, err)
		}
		_bookID = &val
	}
	query := bson.M{
		"_id": _id,
	}
	updateDoc := bson.D{
		{"$set",
			bson.D{
				{"bookingId", _bookID},
			},
		},
	}
	_, err = bs.users.UpdateOne(ctx, query, updateDoc)
	if err != nil {
		return fmt.Errorf("could not update user's bookingId %s", userID)
	}
	return nil
}

func (bs Booking) DeleteUser(ctx context.Context, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("could not parse object id %s: %v", id, err)
	}
	query := bson.M{
		"_id": _id,
	}
	_, err = bs.users.DeleteOne(ctx, query)
	if err != nil {
		return fmt.Errorf("could not delete a user %s", id)
	}
	return nil
}
