package inmem

import (
	"errors"

	"github.com/google/uuid"
	"github.com/picolloo/auth-playground/entities"
)

type UserRepository struct {
	users []*entities.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*entities.User, 0),
	}
}

func (ur *UserRepository) CreateUser(name, email, password string, admin bool) *entities.User {
	return nil
}

func (ur *UserRepository) PromoteUserToAdmin(id uuid.UUID) *entities.User {
	return nil
}

func (ur *UserRepository) GetUsers() []*entities.User {
	return ur.users
}

func (ur *UserRepository) GetUser(id uuid.UUID) (*entities.User, error) {
	for _, user := range ur.users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("Invalid user ID")
}
