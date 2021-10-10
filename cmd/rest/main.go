package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/picolloo/colonia-z33/internals/user"
	"github.com/picolloo/colonia-z33/internals/user/adapters"
	"github.com/picolloo/colonia-z33/internals/user/ports"
)

func main() {
	r := mux.NewRouter()

	user_repo := adapters.NewInmemRepository()
	user_service := user.NewService(user_repo)

	r.Handle("/users", ports.GetUsers(user_service)).Methods("GET")
	r.Handle("/users", ports.CreateUser(user_service)).Methods("POST")

	http.Handle("/", r)

	http.ListenAndServe(":3000", http.DefaultServeMux)
}
