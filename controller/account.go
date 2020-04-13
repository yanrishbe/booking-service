package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

type accountRouter struct {
	*mux.Router
	service Account
}

func newAccountRouter(service Account) *accountRouter {
	router := accountRouter{
		mux.NewRouter().PathPrefix(accountsRoute).Subrouter(),
		service,
	}

	router.Path("").Methods(http.MethodPost).HandlerFunc(router.createAccount)
	router.Path("/{id}").Methods(http.MethodGet).HandlerFunc(router.getAccount)
	router.Path("/{id}").Methods(http.MethodPost).HandlerFunc(router.updateAccount)

	return &router
}

func (ar accountRouter) createAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	validateRights(r.Context(), w, id)

	var a model.Account
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		util.JSONError(http.StatusUnprocessableEntity, w, err)
		return
	}
	response, err := ar.service.Create(r.Context(), a, id)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	util.JSON(w, response)
}

func (ar accountRouter) getAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	validateRights(r.Context(), w, id)
	response, err := ar.service.Get(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	util.JSON(w, response)
}

func (ar accountRouter) updateAccount(w http.ResponseWriter, r *http.Request) {
}
