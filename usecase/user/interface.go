package user

import (
	"github.com/google/uuid"
	"github.com/picolloo/auth-playground/entities"
)

type Writer interface {
	CreateUser(*entities.User) *entities.User
	PromoteUserToAdmin(uuid.UUID) *entities.User
}

type Reader interface {
	GetUsers() []*entities.User
	GetUser(uuid.UUID) *entities.User
	IsEmailInUse(string) bool
}

type Repository interface {
	Reader
	Writer
}
