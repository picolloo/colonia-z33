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
	birth time.Time,
	phone time.Time,
	email string,
	cellPhone time.Time,
	status Status,
	createdAt time.Time,
	nationality string,
	fatherName string,
	motherName string,
	scholarity string,
	category Category,
	RG,
	PIS,
	CPF,
	NIT,
	CEI,
	electorTitle string,
) (*Customer, error) {
	c := NewCustomer(
		name,
		birth,
		phone,
		email,
		cellPhone,
		status,
		createdAt,
		nationality,
		fatherName,
		motherName,
		scholarity,
		category,
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
