package controller

import (
	"net/http"

	"github.com/gorilla/mux"
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
}

func (ar accountRouter) getAccount(w http.ResponseWriter, r *http.Request) {
}

func (ar accountRouter) updateAccount(w http.ResponseWriter, r *http.Request) {
}
