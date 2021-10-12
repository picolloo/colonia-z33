package customer_test

import (
	"testing"
	"time"

	"github.com/picolloo/colonia-z33/internals/customer"
	"github.com/picolloo/colonia-z33/internals/customer/adapters"
)

func TestCreateCustomerWithValidProps(t *testing.T) {

	r := adapters.NewInmemRepository()
	s := customer.NewService(r)

	timestamp := time.Now()
	c, err := s.CreateCustomer(
		"",
		"",
		"",
		"",
		"",
		"Ativo",
		timestamp,
		"",
		"",
		"",
		"",
		"Aposentado Pesca",
		"",
		"",
		"",
		"",
		"",
		"",
	)

	if err != nil {
		t.Fatalf("Expected err: nil Actual: %s", err.Error())
	}

	if c.Name != "" {
		t.Errorf("Expected name: %s Actual: %s", "", c.Name)
	}
	if c.Birth != "" {
		t.Errorf("Expected birth: %s Actual: %s", "", c.Birth)
	}
	if c.Phone != "" {
		t.Errorf("Expected Phone: %s Actual: %s", "", c.Phone)
	}
	if c.Email != "" {
		t.Errorf("Expected Email: %s Actual: %s", "", c.Email)
	}
	if c.CellPhone != "" {
		t.Errorf("Expected CellPhone: %s Actual: %s", "", c.CellPhone)
	}
	if c.Status != customer.Active {
		t.Errorf("Expected Status: %s Actual: %s", customer.Active.String(), c.Status.String())
	}
	if c.CreatedAt != timestamp {
		t.Errorf("Expected CreatedAt: %s Actual: %s", timestamp, c.CreatedAt)
	}
	if c.Nationality != "" {
		t.Errorf("Expected Nationality: %s Actual: %s", "", c.Nationality)
	}
	if c.FatherName != "" {
		t.Errorf("Expected FatherName: %s Actual: %s", "", c.FatherName)
	}
	if c.MotherName != "" {
		t.Errorf("Expected MotherName: %s Actual: %s", "", c.MotherName)
	}
	if c.Scholarity != "" {
		t.Errorf("Expected Scholarity: %s Actual: %s", "", c.Scholarity)
	}
	if c.Category != customer.RetiredFishing {
		t.Errorf("Expected Category: %s Actual: %s", customer.RetiredFishing.String(), c.Category.String())
	}
	if c.PIS != "" {
		t.Errorf("Expected PIS: %s Actual: %s", "", c.PIS)
	}
	if c.CPF != "" {
		t.Errorf("Expected CPF: %s Actual: %s", "", c.CPF)
	}
	if c.NIT != "" {
		t.Errorf("Expected NIT: %s Actual: %s", "", c.NIT)
	}
	if c.CEI != "" {
		t.Errorf("Expected CEI: %s Actual: %s", "", c.CEI)
	}
	if c.ElectorTitle != "" {
		t.Errorf("Expected ElectorTitle: %s Actual: %s", "", c.ElectorTitle)
	}
}

func TestGetCustomers(t *testing.T) {
	r := adapters.NewInmemRepository()
	s := customer.NewService(r)

	timestamp := time.Now()
	c, _ := s.CreateCustomer(
		"",
		"",
		"",
		"",
		"",
		"Ativo",
		timestamp,
		"",
		"",
		"",
		"",
		"Aposentado Pesca",
		"",
		"",
		"",
		"",
		"",
		"",
	)

	cs := s.GetCustomers()
	if len(cs) != 1 {
		t.Fatalf("Expect customer lenght: 1 Actual %d", len(cs))
	}

	if cs[0] != c {
		t.Fatalf("Expect customer: %+v Actual %+v", c, cs[0])
	}
}
