package customer

import "github.com/google/uuid"

type Writer interface {
	CreateCustomer(*Customer) *Customer
	UpdateCustomer(*Customer) *Customer
}

type Reader interface {
	GetCustomers() []*Customer
	GetCustomer(uuid.UUID) *Customer
}

type Repository interface {
	Reader
	Writer
}
