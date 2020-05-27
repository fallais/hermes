package models

import (
	"testing"
	"time"
)

func TestGetAge(t *testing.T) {
	// Create the contact
	contact, err := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/1930")
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}

	// Check the age
	if contact.GetAge() != 90 {
		t.Errorf("The age is incorrect : %d. It should be : %d.", contact.GetAge(), 89)
		t.Fail()
	}

	// Create the contact
	contact, err = NewContact("John", "Doe", "Jojo", "Best friend", "01/01")
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}

	// Check the age
	if contact.GetAge() != 0 {
		t.Errorf("The age is incorrect : %d. It should be : %d.", contact.GetAge(), 0)
		t.Fail()
	}
}

func TestIsBirthdayToday(t *testing.T) {
	// Create the contact
	contact, _ := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/1930")

	// Check the age
	if contact.IsBirthdayToday() {
		t.Errorf("Should not be the birthday")
		t.Fail()
	}

	// Create the contact
	contact, _ = NewContact("John", "Doe", "Jojo", "Best friend", time.Now().Format("02/01/2006"))

	// Check the age
	if !contact.IsBirthdayToday() {
		t.Errorf("Should be the birthday")
		t.Fail()
	}
}

func TestIsBornOnLeapYear(t *testing.T) {
	// Create the contact
	contact, _ := NewContact("John", "Doe", "Jojo", "Best friend", "29/02/1950")

	// Check the leap year
	if !contact.IsBornOnLeapYear() {
		t.Errorf("Should be true")
		t.Fail()
	}

	// Create the contact
	contact, _ = NewContact("John", "Doe", "Jojo", "Best friend", "29/01/1950")

	// Check the leap year
	if contact.IsBornOnLeapYear() {
		t.Errorf("Should be false")
		t.Fail()
	}
}
