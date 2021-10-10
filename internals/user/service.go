package user

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUser(id uuid.UUID) (*User, error) {
	user := s.repo.GetUser(id)

	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (s *Service) GetUsers() []*User {
	return s.repo.GetUsers()
}

func (s *Service) CreateUser(name, email, password string) (*User, error) {
	if invalid_email := s.repo.IsEmailInUse(email); invalid_email {
		return nil, errors.New("Email already in use.")
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	hash_password := string(bytes)

	user := &User{
		Id:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: hash_password,
	}

	s.repo.CreateUser(user)
	return user, nil
}
