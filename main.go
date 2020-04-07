package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/yanrishbe/booking-service/controller"
	"github.com/yanrishbe/booking-service/mongo"
	"github.com/yanrishbe/booking-service/service"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
}

func main() {
	ctx := context.Background()
	// connect to db
	db, err := mongo.NewBooking(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// create service layer
	userService := service.NewUser(db)
	bookingService := service.NewBooking(db)
	accountService := service.NewAccount(db)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-stop
		err := db.Client().Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("gracefully shutting down the server")
		os.Exit(1)
	}()

	log.Fatalln(http.ListenAndServe(":9999", controller.NewRouter(*userService, *accountService, *bookingService)))
}
