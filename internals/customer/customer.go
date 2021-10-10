package customer

import (
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	Active Status = iota
	Inactive
	Disabled
)

type Category int

const (
	RetiredGeneral Category = iota
	RetiredFishing
	Fishing
	Entail
)

type Customer struct {
	Id           uuid.UUID
	Name         string
	Birth        time.Time
	Phone        time.Time
	Celphone     time.Time
	Status       Status
	CreatedAt    time.Time
	Nationality  *string
	FatherName   *string
	MotherName   *string
	Scholarity   *string
	Category     *Category
	RG           *string
	PIS          *string
	CPF          *string
	NIT          *string
	CEI          *string
	ElectorTitle *string
}

func NewCustomer(
	Name string,
	Birth time.Time,
	Phone time.Time,
	Celphone time.Time,
	Status Status,
	CreatedAt time.Time,
	Nationality *string,
	FatherName *string,
	MotherName *string,
	Scholarity *string,
	Category *Category,
	ElectorTitle *string,
) *Customer {
	return &Customer{
		Id:          uuid.New(),
		Name:        Name,
		Birth:       Birth,
		Phone:       Phone,
		Celphone:    Celphone,
		Status:      Status,
		CreatedAt:   CreatedAt,
		Nationality: Nationality,
		FatherName:  FatherName,
		MotherName:  MotherName,
		Scholarity:  Scholarity,
		Category:    Category,
	}
}

func (c *Customer) WithRG(rg string) *Customer {
	c.RG = &rg
	return c
}

func (c *Customer) WithPIS(pis string) *Customer {
	c.PIS = &pis
	return c
}

func (c *Customer) WithCPF(cpf string) *Customer {
	c.CPF = &cpf
	return c
}

func (c *Customer) WithNIT(nit string) *Customer {
	c.NIT = &nit
	return c
}

func (c *Customer) WithCEI(cei string) *Customer {
	c.CEI = &cei
	return c
}
