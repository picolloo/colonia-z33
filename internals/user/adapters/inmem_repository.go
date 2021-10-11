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

func (r *InmemRepository) CreateUser(user *user.User) *user.User {
	r.users = append(r.users, user)
	return user
}

func (r *InmemRepository) GetUsers() []*user.User {
	return r.users
}

func (r *InmemRepository) GetUser(id uuid.UUID) *user.User {
	for _, u := range r.users {
		if u.Id == id {
			return u
		}
	}
	return nil
}

func (ur *InmemRepository) IsEmailInUse(email string) bool {
	for _, u := range ur.users {
		if u.Email == email {
			return true
		}
	}
	return false
}
