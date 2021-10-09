package inmem

import (
	"github.com/google/uuid"
	"github.com/picolloo/colonia-z33/entities"
)

type UserRepository struct {
	users []*entities.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*entities.User, 0),
	}
}

func (ur *UserRepository) CreateUser(user *entities.User) *entities.User {
	ur.users = append(ur.users, user)
	return user
}

func (ur *UserRepository) GetUsers() []*entities.User {
	return ur.users
}

func (ur *UserRepository) GetUser(id uuid.UUID) *entities.User {
	for _, user := range ur.users {
		if user.Id == id {
			return user
		}
	}
	return nil
}

func (ur *UserRepository) IsEmailInUse(email string) bool {
	for _, user := range ur.users {
		if user.Email == email {
			return true
		}
	}
	return false
}
