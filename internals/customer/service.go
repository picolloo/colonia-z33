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
	Name string,
	Birth time.Time,
	Phone time.Time,
	Celphone time.Time,
	Status Status,
	CreatedAt time.Time,
	Nationality *string,
	FatherName *string,
	MotherName *string,
	Scholarity *string,
	Category *Category,
	ElectorTitle *string,
) (*Customer, error) {

	c := NewCustomer(
		Name,
		Birth,
		Phone,
		Celphone,
		Status,
		CreatedAt,
		Nationality,
		FatherName,
		MotherName,
		Scholarity,
		Category,
		ElectorTitle,
	)

	customer := s.repo.CreateCustomer(c)
	return customer, nil
}
