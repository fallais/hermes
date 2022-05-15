package models

import (
	"testing"
	"time"
)

func TestContact(t *testing.T) {
	// Create the contact with a bad year
	_, err := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/193")
	if err == nil {
		t.Fatalf("should err: %s", err)
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
		t.Fatalf("should not err: %s", err)
	}
}

func TestGetAge(t *testing.T) {
	// Create the contact with a year and check age
	contact, err := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/1930")
	if err != nil {
		t.Fatalf("should not err: %s", err)
	}
	if contact.GetAge() != 92 {
		t.Errorf("The age is incorrect : %d. It should be : %d.", contact.GetAge(), 89)
		t.Fail()
	}

	// Create the contact without year and check age
	contact, err = NewContact("John", "Doe", "Jojo", "Best friend", "01/01")
	if err != nil {
		t.Fatalf("should not err: %s", err)
	}
	if contact.GetAge() != 0 {
		t.Errorf("The age is incorrect : %d. It should be : %d.", contact.GetAge(), 0)
		t.Fail()
	}

	// Create the contact with a bad year
	_, err = NewContact("John", "Doe", "Jojo", "Best friend", "01/01/193")
	if err == nil {
		t.Errorf("should err: %s", err)
		t.Fail()
	}
}

func TestIsBirthdayToday(t *testing.T) {
	// Create the contact
	contact, err := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/1930")
	if err != nil {
		t.Fatalf("should not err: %s", err)
	}

	// Check the age
	if contact.IsBirthdayToday() {
		t.Errorf("Should not be the birthday")
		t.Fail()
	}

	// Create the contact
	contact, err = NewContact("John", "Doe", "Jojo", "Best friend", time.Now().Format("02/01/2006"))
	if err != nil {
		t.Fatalf("should not err: %s", err)
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
		t.Fatalf("should not err: %s", err)
	}

	// Check the leap year
	if !contact.IsBornOnLeapYear() {
		t.Errorf("Should be true")
		t.Fail()
	}

	// Create the contact
	contact, err = NewContact("John", "Doe", "Jojo", "Best friend", "29/01/1950")
	if err != nil {
		t.Fatalf("should not err: %s", err)
	}

	// Check the leap year
	if contact.IsBornOnLeapYear() {
		t.Errorf("Should be false")
		t.Fail()
	}
}

func TestGetName(t *testing.T) {
	c, err := NewContact("John", "Doe", "Jojo", "Best friend", "01/01/1933")
	if err != nil {
		t.Fatalf("should not err: %s", err)
	}

	if c.GetName() != "Jojo (John Doe)" {
		t.Errorf("should be Jojo (John Doe) but it is %s", c.GetName())
		t.Fail()
	}

	c, err = NewContact("John", "", "Jojo", "Best friend", "01/01/1933")
	if err != nil {
		t.Fatalf("should not err: %s", err)
	}

	if c.GetName() != "Jojo (John)" {
		t.Errorf("should be Jojo (John) but it is %s", c.GetName())
		t.Fail()
	}

	c, err = NewContact("John", "", "", "Best friend", "01/01/1933")
	if err != nil {
		t.Fatalf("should not err: %s", err)
	}

	if c.GetName() != "John" {
		t.Errorf("should be John but it is %s", c.GetName())
		t.Fail()
	}

	c, err = NewContact("John", "Doe", "", "Best friend", "01/01/1933")
	if err != nil {
		t.Fatalf("should not err: %s", err)
	}

	if c.GetName() != "John Doe" {
		t.Errorf("should be John but it is %s", c.GetName())
		t.Fail()
	}
}
