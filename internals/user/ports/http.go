package ports

import (
	"encoding/json"
	"net/http"

	"github.com/picolloo/colonia-z33/cmd/rest/serializer"
	"github.com/picolloo/colonia-z33/internals/user"
	"github.com/thedevsaddam/govalidator"
)

func GetUsers(s *user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := s.GetUsers()

		response := []*serializer.User{}
		for _, u := range users {
			response = append(response, &serializer.User{
				Id:    u.Id,
				Name:  u.Name,
				Email: u.Email,
			})
		}
		json.NewEncoder(w).Encode(response)
	})
}

func CreateUser(s *user.Service) http.Handler {
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
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(e)
			return
		}

		user, err := s.CreateUser(body.Name, body.Email, body.Password)

		if err != nil {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(e)
			return
		}

		response := &serializer.User{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}

		json.NewEncoder(w).Encode(response)
	})
}
