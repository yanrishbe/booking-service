package main

import (
	"log"
	"net/http"

	"github.com/yanrishbe/booking-service/controller"
)

func main() {
	log.Fatal(http.ListenAndServe(":12345", controller.New()))
}
