package customer

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetCustomer(id uuid.UUID) (*Customer, error) {
	customer := s.repo.GetCustomer(id)

	if customer == nil {
		return nil, errors.New("Customer not found")
	}

	return customer, nil
}

func (s *Service) GetCustomers() []*Customer {
	return s.repo.GetCustomers()
}

func (s *Service) CreateCustomer(
	name string,
	birth string,
	phone string,
	email string,
	cellPhone string,
	status string,
	createdAt time.Time,
	nationality string,
	fatherName string,
	motherName string,
	scholarity string,
	category string,
	RG,
	PIS,
	CPF,
	NIT,
	CEI,
	electorTitle string,
) (*Customer, error) {
	st := StatusFromString(status)
	ct := CategoryFromString(category)

	c := NewCustomer(
		name,
		birth,
		phone,
		email,
		cellPhone,
		st,
		createdAt,
		nationality,
		fatherName,
		motherName,
		scholarity,
		ct,
		RG,
		PIS,
		CPF,
		NIT,
		CEI,
		electorTitle,
	)
	customer := s.repo.CreateCustomer(c)
	return customer, nil
}
