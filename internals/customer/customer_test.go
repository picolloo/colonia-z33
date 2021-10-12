package customer

import (
	"testing"
	"time"
)

func TestCreateCustomeWithAllProps(t *testing.T) {
	timestamp := time.Now()
	c := NewCustomer(
		"",
		"",
		"",
		"",
		"",
		Active,
		timestamp,
		"",
		"",
		"",
		"",
		RetiredFishing,
		"",
		"",
		"",
		"",
		"",
		"",
	)

	if c.Name != "" {
		t.Errorf("Expected name: %s, received: %s", "", c.Name)
	}
	if c.Birth != "" {
		t.Errorf("Expected birth: %s, received: %s", "", c.Birth)
	}
	if c.Phone != "" {
		t.Errorf("Expected Phone: %s, received: %s", "", c.Phone)
	}
	if c.Email != "" {
		t.Errorf("Expected Email: %s, received: %s", "", c.Email)
	}
	if c.CellPhone != "" {
		t.Errorf("Expected CellPhone: %s, received: %s", "", c.CellPhone)
	}
	if c.Status != Active {
		t.Errorf("Expected Status: %s, received: %s", Active.String(), c.Status.String())
	}
	if c.CreatedAt != timestamp {
		t.Errorf("Expected CreatedAt: %s, received: %s", timestamp, c.CreatedAt)
	}
	if c.Nationality != "" {
		t.Errorf("Expected Nationality: %s, received: %s", "", c.Nationality)
	}
	if c.FatherName != "" {
		t.Errorf("Expected FatherName: %s, received: %s", "", c.FatherName)
	}
	if c.MotherName != "" {
		t.Errorf("Expected MotherName: %s, received: %s", "", c.MotherName)
	}
	if c.Scholarity != "" {
		t.Errorf("Expected Scholarity: %s, received: %s", "", c.Scholarity)
	}
	if c.Category != RetiredFishing {
		t.Errorf("Expected Category: %s, received: %s", RetiredFishing.String(), c.Category.String())
	}
	if c.PIS != "" {
		t.Errorf("Expected PIS: %s, received: %s", "", c.PIS)
	}
	if c.CPF != "" {
		t.Errorf("Expected CPF: %s, received: %s", "", c.CPF)
	}
	if c.NIT != "" {
		t.Errorf("Expected NIT: %s, received: %s", "", c.NIT)
	}
	if c.CEI != "" {
		t.Errorf("Expected CEI: %s, received: %s", "", c.CEI)
	}
	if c.ElectorTitle != "" {
		t.Errorf("Expected ElectorTitle: %s, received: %s", "", c.ElectorTitle)
	}
}

func TestCategoryToString(t *testing.T) {
	c := RetiredFishing
	if c.String() != "Aposentado Pesca" {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", "Aposentado Pesca", c.String())
	}

	c = RetiredGeneral
	if c.String() != "Aposentado Geral" {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", "Aposentado Geral", c.String())
	}

	c = Entail
	if c.String() != "Vínculo" {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", "Vínculo", c.String())
	}

	c = Fishing
	if c.String() != "Pesca Exclusiva" {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", "Pesca Exclusiva", c.String())
	}

	c = Unknown
	if c.String() != "Desconhecida" {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", "Desconhecida", c.String())
	}
}

func TestStringToCategory(t *testing.T) {
	c := "Aposentado Pesca"
	if CategoryFromString(c) != RetiredFishing {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", RetiredFishing, c)
	}

	c = "Aposentado Geral"
	if CategoryFromString(c) != RetiredGeneral {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", RetiredGeneral, c)
	}

	c = "Pesca Exclusiva"
	if CategoryFromString(c) != Fishing {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", Fishing, c)
	}

	c = "Vínculo"
	if CategoryFromString(c) != Entail {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", Entail, c)
	}

	c = ""
	if CategoryFromString(c) != Unknown {
		t.Errorf("Invalid category string. Expected: %s, Actual: %s", Unknown, c)
	}

}
