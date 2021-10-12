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

func (r *InmemRepository) CreateCustomer(customer *customer.Customer) *customer.Customer {
	r.customers = append(r.customers, customer)
	return customer
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
