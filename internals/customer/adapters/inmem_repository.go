package adapters

import (
	"github.com/google/uuid"
	"github.com/picolloo/colonia-z33/internals/customer"
)

type InmemRepository struct {
	customers []*customer.Customer
}

func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		customers: make([]*customer.Customer, 0),
	}
}

func (r *InmemRepository) CreateCustomer(user *customer.Customer) *customer.Customer {
	r.customers = append(r.customers, user)
	return user
}

func (r *InmemRepository) GetCustomers() []*customer.Customer {
	return r.customers
}

func (r *InmemRepository) GetCustomer(id uuid.UUID) *customer.Customer {
	for _, c := range r.customers {
		if c.Id == id {
			return c
		}
	}
	return nil
}

func (r *InmemRepository) UpdateCustomer(c *customer.Customer) *customer.Customer {
	for i, c := range r.customers {
		if c.Id == c.Id {
			new_customer := append(r.customers[:i], c)
			r.customers = append(new_customer, r.customers[i+1:]...)
			return c
		}
	}
	return nil
}
