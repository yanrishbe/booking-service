package util

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JSON(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		log.Println(err)
		return
	}
}

func JSONError(code int, w http.ResponseWriter, err error) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		logrus.Debug("encoding error")
	}
}
