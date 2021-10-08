package user

import "github.com/picolloo/auth-playground/entities"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUsers() []*entities.User {
	return s.repo.GetUsers()
}

func (s *Service) CreateUser() (*entities.User, error) {
	return nil, nil
}
