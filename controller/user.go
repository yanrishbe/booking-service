package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

type userRouter struct {
	*mux.Router
	service User
}

func newUserRouter(service User) *userRouter {
	router := userRouter{
		mux.NewRouter().PathPrefix(userRoute).Subrouter(),
		service,
	}

	router.Path("").Methods(http.MethodPost).HandlerFunc(router.createUser)
	router.Path("/{id}").Methods(http.MethodGet).HandlerFunc(router.getUser)
	router.Path("/{id}").Methods(http.MethodPost).HandlerFunc(router.updateUser)
	router.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(router.deleteUser)
	router.Path("").Methods(http.MethodGet).HandlerFunc(router.getAllUsers)

	return &router
}

func (ur userRouter) createUser(w http.ResponseWriter, r *http.Request) {
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
func (ur userRouter) getUser(w http.ResponseWriter, r *http.Request) {
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
func (ur userRouter) updateUser(w http.ResponseWriter, r *http.Request) {
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
func (ur userRouter) deleteUser(w http.ResponseWriter, r *http.Request) {
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
func (ur userRouter) getAllUsers(w http.ResponseWriter, r *http.Request) {
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
