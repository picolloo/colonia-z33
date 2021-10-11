package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/picolloo/colonia-z33/internals/common/middlewares"
	"github.com/picolloo/colonia-z33/internals/customer"
	customer_adapters "github.com/picolloo/colonia-z33/internals/customer/adapters"
	customer_ports "github.com/picolloo/colonia-z33/internals/customer/ports"
	"github.com/picolloo/colonia-z33/internals/user"
	user_adapters "github.com/picolloo/colonia-z33/internals/user/adapters"
	user_ports "github.com/picolloo/colonia-z33/internals/user/ports"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	n := negroni.Classic()
	n.UseHandler(r)
	n.Use(negroni.HandlerFunc(middlewares.JsonMiddleware))

	user_repo := user_adapters.NewInmemRepository()
	user_service := user.NewService(user_repo)
	user_router := r.PathPrefix("/users").Subrouter()
	user_ports.NewRouter(user_service, user_router)

	customer_repo := customer_adapters.NewInmemRepository()
	customer_service := customer.NewService(customer_repo)
	customer_router := r.PathPrefix("/customers").Subrouter()
	customer_ports.NewRouter(customer_service, customer_router)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
