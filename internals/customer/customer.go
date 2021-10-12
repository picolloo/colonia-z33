package customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id           uuid.UUID
	Name         string
	Phone        string
	Email        string
	CellPhone    string
	Nationality  string
	FatherName   string
	MotherName   string
	Scholarity   string
	RG           string
	PIS          string
	CPF          string
	NIT          string
	CEI          string
	ElectorTitle string
	Birth        string
	Status       Status
	Category     Category
	CreatedAt    time.Time
}

func NewCustomer(
	name string,
	birth string,
	phone string,
	email string,
	cellphone string,
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
		CellPhone:    cellphone,
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
	Unknown
)

func (c Category) String() string {
	switch c {
	case RetiredGeneral:
		return "Aposentado Geral"
	case RetiredFishing:
		return "Aposentado Pesca"
	case Fishing:
		return "Pesca Exclusiva"
	case Entail:
		return "Vínculo"
	}
	return "Desconhecida"
}

func CategoryFromString(value string) Category {
	switch value {
	case "Aposentado Geral":
		return RetiredGeneral
	case "Aposentado Pesca":
		return RetiredFishing
	case "Pesca Exclusiva":
		return Fishing
	case "Vínculo":
		return Entail
	}
	return Unknown
}
