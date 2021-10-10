package adapters

import (
	"github.com/google/uuid"
	"github.com/picolloo/colonia-z33/internals/user"
)

type InmemRepository struct {
	users []*user.User
}

func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		users: make([]*user.User, 0),
	}
}

func (ur *InmemRepository) CreateUser(user *user.User) *user.User {
	ur.users = append(ur.users, user)
	return user
}

func (ur *InmemRepository) GetUsers() []*user.User {
	return ur.users
}

func (ur *InmemRepository) GetUser(id uuid.UUID) *user.User {
	for _, user := range ur.users {
		if user.Id == id {
			return user
		}
	}
	return nil
}

func (ur *InmemRepository) IsEmailInUse(email string) bool {
	for _, user := range ur.users {
		if user.Email == email {
			return true
		}
	}
	return false
}
