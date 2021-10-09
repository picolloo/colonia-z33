package user

import (
	"github.com/google/uuid"
	"github.com/picolloo/colonia-z33/entities"
)

type Writer interface {
	CreateUser(*entities.User) *entities.User
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
