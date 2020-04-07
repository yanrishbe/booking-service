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
		mux.NewRouter().PathPrefix(accountRoute).Subrouter(),
		service,
	}

	router.Path("").Methods(http.MethodPost).HandlerFunc(router.createAccount)
	router.Path("/{id}").Methods(http.MethodGet).HandlerFunc(router.getAccount)
	router.Path("/{id}").Methods(http.MethodPost).HandlerFunc(router.updateAccount)

	return &router
}

func (ar accountRouter) createAccount(w http.ResponseWriter, r *http.Request) {
	// pID := mux.Vars(r)
	// data, err := middleware.DataFromContext(r.Context())
	// if err != nil {
	// 	middleware.JSONError(w, e.InvalidMiddlewareContext(err), http.StatusBadRequest)
	// 	return
	// }
	// err = br.service.Delete(r.Context(), data, pID["id"])
	// if err != nil {
	// 	middleware.JSONError(w, err, http.StatusInternalServerError)
	// }
	// middleware.Empty(w, http.StatusCreated)
}

func (ar accountRouter) getAccount(w http.ResponseWriter, r *http.Request) {
	// pID := mux.Vars(r)
	// data, err := middleware.DataFromContext(r.Context())
	// if err != nil {
	// 	middleware.JSONError(w, e.InvalidMiddlewareContext(err), http.StatusBadRequest)
	// 	return
	// }
	// err = ur.service.Delete(r.Context(), data, pID["id"])
	// if err != nil {
	// 	middleware.JSONError(w, err, http.StatusInternalServerError)
	// }
	// middleware.Empty(w, http.StatusCreated)
}

func (ar accountRouter) updateAccount(w http.ResponseWriter, r *http.Request) {
	// pID := mux.Vars(r)
	// data, err := middleware.DataFromContext(r.Context())
	// if err != nil {
	// 	middleware.JSONError(w, e.InvalidMiddlewareContext(err), http.StatusBadRequest)
	// 	return
	// }
	// err = ur.service.Delete(r.Context(), data, pID["id"])
	// if err != nil {
	// 	middleware.JSONError(w, err, http.StatusInternalServerError)
	// }
	// middleware.Empty(w, http.StatusCreated)
}
