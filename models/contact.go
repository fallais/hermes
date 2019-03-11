package models

import (
	"time"

	"github.com/dchest/uniuri"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Contact is a contact.
type Contact struct {
	ID          string
	Firstname   string
	Lastname    string
	Nickname    string
	Description string
	Birthdate   *Birthdate
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewContact returns a valid Contact instance
func NewContact(firstname, lastname, nickname, description string, bd *Birthdate) *Contact {
	return &Contact{
		ID:          uniuri.New(),
		Firstname:   firstname,
		Lastname:    lastname,
		Nickname:    nickname,
		Description: description,
		Birthdate:   bd,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetAge returns the age of the contact. 0 if there is no Year.
func (c *Contact) GetAge() int {
	if c.Birthdate.Year == 0 {
		return 0
	}

	return time.Now().Year() - c.Birthdate.Year
}

// IsBirthdayToday returns true if it is the birthday of the contact.
func (c *Contact) IsBirthdayToday() bool {
	if c.Birthdate.Day == time.Now().Day() && c.Birthdate.Month == int(time.Now().Month()) {
		return true
	}

	return false
}
