package user

import (
	"github.com/google/uuid"
	"github.com/picolloo/auth-playground/entities"
)

type Writer interface {
	CreateUser(name, email, password string, admin bool) *entities.User
	PromoteUserToAdmin(id uuid.UUID) *entities.User
}

type Reader interface {
	GetUsers() []*entities.User
	GetUser(id uuid.UUID) (*entities.User, error)
}

type Repository interface {
	Reader
	Writer
}
