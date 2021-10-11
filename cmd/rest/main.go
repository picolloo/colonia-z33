package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/picolloo/colonia-z33/internals/customer"
	customer_adapters "github.com/picolloo/colonia-z33/internals/customer/adapters"
	customer_ports "github.com/picolloo/colonia-z33/internals/customer/ports"
	"github.com/picolloo/colonia-z33/internals/user"
	user_adapters "github.com/picolloo/colonia-z33/internals/user/adapters"
	user_ports "github.com/picolloo/colonia-z33/internals/user/ports"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)

	user_repo := user_adapters.NewInmemRepository()
	user_service := user.NewService(user_repo)
	user_router := r.PathPrefix("/users").Subrouter()
	user_ports.NewRouter(user_service, user_router)

	customer_repo := customer_adapters.NewInmemRepository()
	customer_service := customer.NewService(customer_repo)
	customer_router := r.PathPrefix("/customers").Subrouter()
	customer_ports.NewRouter(customer_service, customer_router)

	http.Handle("/", r)

	http.ListenAndServe(":3000", http.DefaultServeMux)
}
