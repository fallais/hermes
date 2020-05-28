package models

import (
	"testing"
	"time"
)

func TestContact(t *testing.T) {
	// Create the contact with a bad year
	_, err := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/193")
	if err == nil {
		t.Errorf("should err: %s", err)
		t.Fail()
	}

	// Create the contact with a bad firstname and nickname
	_, err = NewContact("", "Doe", "", "Best friend", "01/01/1995")
	if err == nil {
		t.Errorf("should err: %s", err)
		t.Fail()
	}

	// Create the contact good
	_, err = NewContact("John", "Doe", "", "Best friend", "01/01/1995")
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}
}

func TestGetAge(t *testing.T) {
	// Create the contact with a year and check age
	contact, err := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/1930")
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}
	if contact.GetAge() != 90 {
		t.Errorf("The age is incorrect : %d. It should be : %d.", contact.GetAge(), 89)
		t.Fail()
	}

	// Create the contact without year and check age
	contact, err = NewContact("John", "Doe", "Jojo", "Best friend", "01/01")
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}
	if contact.GetAge() != 0 {
		t.Errorf("The age is incorrect : %d. It should be : %d.", contact.GetAge(), 0)
		t.Fail()
	}

	// Create the contact with a bad year
	contact, err = NewContact("John", "Doe", "Jojo", "Best friend", "01/01/193")
	if err == nil {
		t.Errorf("should err: %s", err)
		t.Fail()
	}
}

func TestIsBirthdayToday(t *testing.T) {
	// Create the contact
	contact, err := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/1930")
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}

	// Check the age
	if contact.IsBirthdayToday() {
		t.Errorf("Should not be the birthday")
		t.Fail()
	}

	// Create the contact
	contact, err = NewContact("John", "Doe", "Jojo", "Best friend", time.Now().Format("02/01/2006"))
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}

	// Check the age
	if !contact.IsBirthdayToday() {
		t.Errorf("Should be the birthday")
		t.Fail()
	}
}

func TestIsBornOnLeapYear(t *testing.T) {
	// Create the contact
	contact, err := NewContact("John", "Doe", "Jojo", "Best friend", "29/02/1952")
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}

	// Check the leap year
	if !contact.IsBornOnLeapYear() {
		t.Errorf("Should be true")
		t.Fail()
	}

	// Create the contact
	contact, err = NewContact("John", "Doe", "Jojo", "Best friend", "29/01/1950")
	if err != nil {
		t.Errorf("should not err: %s", err)
		t.Fail()
	}

	// Check the leap year
	if contact.IsBornOnLeapYear() {
		t.Errorf("Should be false")
		t.Fail()
	}
}
