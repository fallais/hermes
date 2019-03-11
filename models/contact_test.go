package models

import (
	"testing"
	"time"
)

func TestGetAge(t *testing.T) {
	// Create the birthdate
	bd := NewBirthdate(01, 01, 1930)

	// Create the contact
	contact := NewContact("John", "Doe", "Jojo", "Best friend", bd)

	// Check the age
	if contact.GetAge() != 89 {
		t.Errorf("The age is incorrect : %d. It should be : %d.", contact.GetAge(), 89)
		t.Fail()
	}

	// Create the birthdate
	bd = NewBirthdate(01, 01, 0)

	// Create the contact
	contact = NewContact("John", "Doe", "Jojo", "Best friend", bd)

	// Check the age
	if contact.GetAge() != 0 {
		t.Errorf("The age is incorrect : %d. It should be : %d.", contact.GetAge(), 0)
		t.Fail()
	}
}

func TestIsBirthdayToday(t *testing.T) {
	// Create the birthdate
	bd := NewBirthdate(01, 01, 1930)

	// Create the contact
	contact := NewContact("John", "Doe", "Jojo", "Best friend", bd)

	// Check the age
	if contact.IsBirthdayToday() {
		t.Errorf("Should not be the birthday")
		t.Fail()
	}

	// Create the birthdate
	bd = NewBirthdate(time.Now().Day(), int(time.Now().Month()), time.Now().Year())

	// Create the contact
	contact = NewContact("John", "Doe", "Jojo", "Best friend", bd)

	// Check the age
	if !contact.IsBirthdayToday() {
		t.Errorf("Should be the birthday")
		t.Fail()
	}
}
