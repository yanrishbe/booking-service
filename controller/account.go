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

const accountsRoute = "/{id}/accounts"

func newAccountRouter(service Account, userRouter userRouter) *accountRouter {
	router := accountRouter{
		userRouter.PathPrefix("").Subrouter(),
		service,
	}

	router.Path("").Methods(http.MethodPost).HandlerFunc(validateTokenMiddleware(router.createAccount))
	router.Path(accountsRoute + "/{accountId}").Methods(http.MethodGet).HandlerFunc(validateTokenMiddleware(router.getAccount))
	router.Path(accountsRoute + "/{accountId}").Methods(http.MethodPut).HandlerFunc(validateTokenMiddleware(router.updateAccount))

	return &router
}

func (ar accountRouter) createAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}

	var a model.Account
	err = json.NewDecoder(r.Body).Decode(&a)
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
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}
	accountId := mux.Vars(r)["accountId"]
	response, err := ar.service.Get(r.Context(), accountId)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	util.JSON(w, response)
}

func (ar accountRouter) updateAccount(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}

	var a model.Account
	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		util.JSONError(http.StatusUnprocessableEntity, w, err)
		return
	}

	accountId := mux.Vars(r)["accountId"]
	err = ar.service.Update(r.Context(), a, accountId, id)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
