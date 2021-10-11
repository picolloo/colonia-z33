package customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id           uuid.UUID
	Name         string
	Birth        time.Time
	Phone        time.Time
	Email        string
	CellPhone    time.Time
	Status       Status
	CreatedAt    time.Time
	Nationality  string
	FatherName   string
	MotherName   string
	Scholarity   string
	Category     Category
	RG           string
	PIS          string
	CPF          string
	NIT          string
	CEI          string
	ElectorTitle string
}

func NewCustomer(
	name string,
	birth time.Time,
	phone time.Time,
	email string,
	CellPhone time.Time,
	status Status,
	createdAt time.Time,
	nationality string,
	fatherName string,
	motherName string,
	scholarity string,
	category Category,
	RG string,
	PIS string,
	CPF string,
	NIT string,
	CEI string,
	electorTitle string,
) *Customer {
	return &Customer{
		Id:           uuid.New(),
		Name:         name,
		Birth:        birth,
		Phone:        phone,
		Email:        email,
		CellPhone:    CellPhone,
		Status:       status,
		CreatedAt:    createdAt,
		Nationality:  nationality,
		FatherName:   fatherName,
		MotherName:   motherName,
		Scholarity:   scholarity,
		Category:     category,
		PIS:          PIS,
		CPF:          CPF,
		NIT:          NIT,
		CEI:          CEI,
		ElectorTitle: electorTitle,
	}
}

type Status int

const (
	Active Status = iota + 1
	Inactive
	Disabled
)

func (s Status) String() string {
	switch s {
	case Active:
		return "Ativo"
	case Inactive:
		return "Inativo"
	case Disabled:
		return "Desabilitado"
	}
	return ""
}

func StatusFromString(value string) Status {
	switch value {
	case "Ativo":
		return Active
	case "Inativo":
		return Inactive
	case "Desabilitado":
		return Disabled
	}
	return Inactive
}

type Category int

const (
	RetiredGeneral Category = iota + 1
	RetiredFishing
	Fishing
	Entail
)

func (c Category) String() string {
	switch c {
	case RetiredGeneral:
		return "Aposentado Geral"
	case RetiredFishing:
		return "Aposentado Pesca"
	case Fishing:
		return "Pesca Exlusiva"
	case Entail:
		return "Vínculo"
	}
	return ""
}

func CategoryFromString(value string) Category {
	switch value {
	case "Aposentado Geral":
		return RetiredGeneral
	case "Aposentado Pesca":
		return RetiredFishing
	case "Pesca Exlusiva":
		return Fishing
	case "Vínculo":
		return Entail
	}
	return Entail
}
