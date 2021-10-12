package adapters

import (
	"testing"

	"github.com/google/uuid"
	"github.com/picolloo/colonia-z33/internals/customer"
)

func TestCreateCustomer(t *testing.T) {
	repo := NewInmemRepository()

	c := &customer.Customer{}
	repo.CreateCustomer(c)

	if len(repo.customers) != 1 {
		t.Fatalf("Invalid customers. Expect: 1 Actual: %d", len(repo.customers))
	}

	if repo.customers[0] != c {
		t.Fatalf("Expect: %+v Actual: %+v", c, repo.customers[0])
	}
}

func TestGetCustomerValidId(t *testing.T) {
	repo := NewInmemRepository()

	c := &customer.Customer{Id: uuid.New()}
	repo.customers = []*customer.Customer{c}

	nc := repo.GetCustomer(c.Id)
	if nc != c {
		t.Fatalf("Expect: %+v Actual: %+v", c, nc)
	}
}

func TestGetCustomerinvalidId(t *testing.T) {
	repo := NewInmemRepository()

	repo.customers = []*customer.Customer{
		{Id: uuid.New()},
	}

	c := repo.GetCustomer(uuid.New())
	if c != nil {
		t.Fatalf("Expect: nil Actual: %+v", c)
	}
}

func TestGetCustomers(t *testing.T) {
	repo := NewInmemRepository()

	c := &customer.Customer{Id: uuid.New()}
	repo.customers = []*customer.Customer{c}

	cs := repo.GetCustomers()
	if len(cs) != 1 {
		t.Fatalf("Expect: 1 Actual: %d", len(cs))
	}

	if cs[0] != c {
		t.Fatalf("Expect: %+v Actual: %+v", c, cs[0])
	}
}
