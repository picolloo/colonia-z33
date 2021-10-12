package ports

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/picolloo/colonia-z33/internals/customer"
	"github.com/picolloo/colonia-z33/internals/customer/adapters"
	"github.com/thedevsaddam/govalidator"
)

func NewRouter(s *customer.Service, r *mux.Router) *mux.Router {
	r.Handle("", getCustomers(s)).Methods("GET")
	r.Handle("/{id}", getCustomer(s)).Methods("GET")
	r.Handle("", createCustomer(s)).Methods("POST")
	return r
}

func getCustomer(s *customer.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := uuid.Parse(vars["id"])

		if err != nil {
			http.Error(w, "Invalid ID.", http.StatusBadRequest)
			return
		}
		customer, err := s.GetCustomer(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		response := &adapters.Customer{
			Id:           customer.Id,
			Name:         customer.Name,
			Email:        customer.Email,
			Birth:        customer.Birth,
			Phone:        customer.Phone,
			CellPhone:    customer.CellPhone,
			Status:       customer.Status.String(),
			CreatedAt:    customer.CreatedAt,
			Nationality:  customer.Nationality,
			FatherName:   customer.FatherName,
			MotherName:   customer.MotherName,
			Scholarity:   customer.Scholarity,
			Category:     customer.Category.String(),
			RG:           customer.RG,
			PIS:          customer.PIS,
			CPF:          customer.CPF,
			NIT:          customer.NIT,
			CEI:          customer.CEI,
			ElectorTitle: customer.ElectorTitle,
		}

		json.NewEncoder(w).Encode(response)
	})
}

func getCustomers(s *customer.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customers := s.GetCustomers()

		response := []*adapters.Customer{}
		for _, u := range customers {
			response = append(response, &adapters.Customer{
				Id:           u.Id,
				Name:         u.Name,
				Email:        u.Email,
				Birth:        u.Birth,
				Phone:        u.Phone,
				CellPhone:    u.CellPhone,
				Status:       u.Status.String(),
				CreatedAt:    u.CreatedAt,
				Nationality:  u.Nationality,
				FatherName:   u.FatherName,
				MotherName:   u.MotherName,
				Scholarity:   u.Scholarity,
				Category:     u.Category.String(),
				RG:           u.RG,
				PIS:          u.PIS,
				CPF:          u.CPF,
				NIT:          u.NIT,
				CEI:          u.CEI,
				ElectorTitle: u.ElectorTitle,
			})
		}
		json.NewEncoder(w).Encode(response)
	})
}

func createCustomer(s *customer.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body adapters.Customer

		rules := govalidator.MapData{
			"name":         []string{},
			"email":        []string{"email"},
			"birth":        []string{"date"},
			"phone":        []string{"date"},
			"cellPhone":    []string{"date"},
			"status":       []string{"in:Ativo,Inativo,Desabilitado"},
			"createdAt":    []string{"date"},
			"nationality":  []string{},
			"fatherName":   []string{},
			"motherName":   []string{},
			"scholarity":   []string{},
			"category":     []string{"in:Aposentado Geral,Aposentado Pesca,Pesca Exclusiva,Vínculo"},
			"RG":           []string{},
			"PIS":          []string{},
			"CPF":          []string{},
			"NIT":          []string{},
			"CEI":          []string{},
			"electorTitle": []string{},
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

		customer, err := s.CreateCustomer(
			body.Name,
			body.Birth,
			body.Phone,
			body.Email,
			body.CellPhone,
			body.Status,
			body.CreatedAt,
			body.Nationality,
			body.FatherName,
			body.MotherName,
			body.Scholarity,
			body.Category,
			body.RG,
			body.PIS,
			body.CPF,
			body.NIT,
			body.CEI,
			body.ElectorTitle,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := &adapters.Customer{
			Id:           customer.Id,
			Name:         customer.Name,
			Email:        customer.Email,
			Birth:        customer.Birth,
			Phone:        customer.Phone,
			CellPhone:    customer.CellPhone,
			Status:       customer.Status.String(),
			CreatedAt:    customer.CreatedAt,
			Nationality:  customer.Nationality,
			FatherName:   customer.FatherName,
			MotherName:   customer.MotherName,
			Scholarity:   customer.Scholarity,
			Category:     customer.Category.String(),
			RG:           customer.RG,
			PIS:          customer.PIS,
			CPF:          customer.CPF,
			NIT:          customer.NIT,
			CEI:          customer.CEI,
			ElectorTitle: customer.ElectorTitle,
		}

		json.NewEncoder(w).Encode(response)
	})
}
