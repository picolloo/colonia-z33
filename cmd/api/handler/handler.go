package handler

import (
	"encoding/json"
	"net/http"

	"github.com/picolloo/auth-playground/cmd/api/serializer"
	"github.com/picolloo/auth-playground/usecase/user"
	"github.com/thedevsaddam/govalidator"
)

func GetUsers(s *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users := s.GetUsers()

		response := make([]*serializer.User, len(users))
		for _, u := range users {
			response = append(response, &serializer.User{
				Id:    u.Id,
				Name:  u.Name,
				Email: u.Email,
				Admin: u.Admin,
			})
		}
		data, _ := json.Marshal(response)
		w.Write(data)
	}
}

func CreateUser(s *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Admin    bool   `json:"admin"`
			Password string `json:"password"`
		}{}

		rules := govalidator.MapData{
			"name":     []string{"required"},
			"email":    []string{"required", "email"},
			"password": []string{"required"},
			"admin":    []string{"bool"},
		}

		opts := govalidator.Options{
			Request:         r,
			Data:            &body,
			Rules:           rules,
			RequiredDefault: false,
		}
		err := govalidator.New(opts).ValidateJSON()

		if len(err) > 0 {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(body)
	}
}
