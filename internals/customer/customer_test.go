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
	if c.Status != Active {
		t.Errorf("Expected Status: %s Actual: %s", Active.String(), c.Status.String())
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
	if c.Category != RetiredFishing {
		t.Errorf("Expected Category: %s Actual: %s", RetiredFishing.String(), c.Category.String())
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
