package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		log.Println("encoding error")
		return
	}
}

func JSONError(code int, w http.ResponseWriter, err error) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		log.Println("encoding error")
	}
}
