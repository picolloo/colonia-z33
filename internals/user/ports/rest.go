package ports

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/picolloo/colonia-z33/internals/user"
	"github.com/picolloo/colonia-z33/internals/user/adapters"
	"github.com/thedevsaddam/govalidator"
)

func NewRouter(s *user.Service, r *mux.Router) *mux.Router {
	r.Handle("", getUsers(s)).Methods("GET")
	r.Handle("", createUser(s)).Methods("POST")
	return r
}

func getUsers(s *user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := s.GetUsers()

		response := []*adapters.User{}
		for _, u := range users {
			response = append(response, &adapters.User{
				Id:    u.Id,
				Name:  u.Name,
				Email: u.Email,
			})
		}
		json.NewEncoder(w).Encode(response)
	})
}

func createUser(s *user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}

		rules := govalidator.MapData{
			"name":     []string{"required"},
			"email":    []string{"required", "email"},
			"password": []string{"required"},
		}

		opts := govalidator.Options{
			Request:         r,
			Data:            &body,
			Rules:           rules,
			RequiredDefault: false,
		}
		e := govalidator.New(opts).ValidateJSON()

		if len(e) > 0 {
			response, _ := json.Marshal(e)
			http.Error(w, string(response), http.StatusBadRequest)
			return
		}

		user, err := s.CreateUser(body.Name, body.Email, body.Password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := &adapters.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		json.NewEncoder(w).Encode(response)
	})
}
