package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/picolloo/auth-playground/entities"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUser(id uuid.UUID) (*entities.User, error) {
	user := s.repo.GetUser(id)

	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (s *Service) GetUsers() []*entities.User {
	return s.repo.GetUsers()
}

func (s *Service) CreateUser(name, email, password string, admin bool) (*entities.User, error) {
	user := &entities.User{
		Id:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: password,
		Admin:    admin,
	}

	s.repo.CreateUser(user)
	return user, nil
}
