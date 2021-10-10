package user

import (
	"github.com/google/uuid"
)

type Writer interface {
	CreateUser(*User) *User
}

type Reader interface {
	GetUsers() []*User
	GetUser(uuid.UUID) *User
	IsEmailInUse(string) bool
}

type Repository interface {
	Reader
	Writer
}
