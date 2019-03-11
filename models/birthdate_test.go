package models

import (
	"testing"
)

func TestToString(t *testing.T) {
	// Create the birthdate
	bd := NewBirthdate(01, 01, 1930)

	// Check the string
	if bd.ToString() != "01/01/1930" {
		t.Errorf("The string representation of the birthdate is incorrect : %s. It should be : %s.", bd.ToString(), "01/01/1930")
		t.Fail()
	}

	// Create the birthdate
	bd = NewBirthdate(01, 01, 0)

	// Check the string
	if bd.ToString() != "01/01" {
		t.Errorf("The string representation of the birthdate is incorrect : %s. It should be : %s.", bd.ToString(), "01/01")
		t.Fail()
	}
}