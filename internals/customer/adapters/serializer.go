package adapters

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Birth        string    `json:"birth"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	CellPhone    string    `json:"cellPhone"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	Nationality  string    `json:"nationality"`
	FatherName   string    `json:"fatherName"`
	MotherName   string    `json:"motherName"`
	Scholarity   string    `json:"scholarity"`
	Category     string    `json:"category"`
	RG           string    `json:"RG"`
	PIS          string    `json:"PIS"`
	CPF          string    `json:"CPF"`
	NIT          string    `json:"NIT"`
	CEI          string    `json:"CEI"`
	ElectorTitle string    `json:"electorTitle"`
}
