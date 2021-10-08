package entities

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Name     string
	Email    string
	Password string
	Admin    bool
}

func NewUser(name, email, password string, admin bool) *User {
	return &User{
		Id:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: password,
		Admin:    admin,
	}
}
