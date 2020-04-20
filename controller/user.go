package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yanrishbe/booking-service/model"
	"github.com/yanrishbe/booking-service/util"
)

type userRouter struct {
	*mux.Router
	service User
}

func newUserRouter(service User) *userRouter {
	router := userRouter{
		mux.NewRouter().PathPrefix(usersRoute).Subrouter(),
		service,
	}

	router.Path("").Methods(http.MethodPost).HandlerFunc(router.createUser)
	router.Path("/login").Methods(http.MethodPost).HandlerFunc(router.loginUser)
	router.Path("/{id}").Methods(http.MethodGet).HandlerFunc(validateTokenMiddleware(router.getUser))
	router.Path("/{id}").Methods(http.MethodPut).HandlerFunc(validateTokenMiddleware(router.updateUser))
	router.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(validateTokenMiddleware(router.deleteUser))

	return &router
}

func (ur userRouter) createUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		util.JSONError(http.StatusUnprocessableEntity, w, err)
		return
	}
	id, err := ur.service.Create(r.Context(), u)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	util.JSON(w, id)
}

func (ur userRouter) loginUser(w http.ResponseWriter, r *http.Request) {
	var l util.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		util.JSONError(http.StatusUnprocessableEntity, w, err)
		return
	}
	id, err := ur.service.Login(r.Context(), l)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}
	token, err := createToken(l.Email, id)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	util.JSON(w, token)
}

func (ur userRouter) getUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}
	response, err := ur.service.Get(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	util.JSON(w, response)
}

func (ur userRouter) updateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := validateRights(r.Context(), id)
	if err != nil {
		util.JSONError(http.StatusUnauthorized, w, err)
		return
	}
	var u util.UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		util.JSONError(http.StatusUnprocessableEntity, w, err)
		return
	}
	u.ID = id
	err = ur.service.Update(r.Context(), u)
	if err != nil {
		util.JSONError(http.StatusInternalServerError, w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (ur userRouter) deleteUser(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["id"]
	// err := validateRights(r.Context(), id)
	// if err != nil {
	// 	util.JSONError(http.StatusUnauthorized, w, err)
	// 	return
	// }
	// err = ur.service.Delete(r.Context(), id)
	// if err != nil {
	// 	util.JSONError(http.StatusInternalServerError, w, err)
	// 	return
	// }
	// w.WriteHeader(http.StatusNoContent)
}

func validateRights(ctx context.Context, id string) error {
	auth, ok := ctx.Value("auth").(Authorization)
	if !ok {
		return fmt.Errorf("user doesn't have enough rights to use resource")
	}
	if auth.ID != id && auth.Role != "admin" {
		return fmt.Errorf("user doesn't have enough rights to use resource")

	}
	return nil
}
