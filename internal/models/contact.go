package models

import (
	"fmt"
	"time"
)

// BirthdateRegex is the Regex used to parse the birthdate.
const BirthdateRegex = "(\\d{2})\\/(\\d{2})(?:(?:\\/)(\\d{4}))?"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Contact is a contact.
type Contact struct {
	Firstname   string
	Lastname    string
	Nickname    string
	Description string
	Birthdate   time.Time
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewContact returns a valid Contact instance
func NewContact(firstname, lastname, nickname, description, birthdate string) (*Contact, error) {
	// Validate the firstname and nickname
	if firstname == "" && nickname == "" {
		return nil, fmt.Errorf("you must provide at least a firstname or a nickname")
	}

	// Create the model
	model := &Contact{
		Firstname:   firstname,
		Lastname:    lastname,
		Nickname:    nickname,
		Description: description,
	}

	// Parse the birthdate
	t, err := time.Parse("02/01/2006", birthdate)
	if err != nil {
		t, err = time.Parse("02/01", birthdate)
		if err != nil {
			return nil, fmt.Errorf("cannot parse the birthdate: %s", err)
		}
	}
	model.Birthdate = t

	return model, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetAge returns the age of the contact. 0 if there is no Year.
func (model *Contact) GetAge() int {
	if model.Birthdate.Year() == 0 {
		return 0
	}

	return time.Now().Year() - model.Birthdate.Year()
}

// IsBirthdayToday returns true if it is the birthday of the contact.
func (model *Contact) IsBirthdayToday() bool {
	if model.Birthdate.Day() == time.Now().Day() && model.Birthdate.Month() == time.Now().Month() {
		return true
	}

	return false
}

// IsBornOnLeapYear returns true if the contact is born on Frebruary the 29th.
func (model *Contact) IsBornOnLeapYear() bool {
	return model.Birthdate.Day() == 29 && model.Birthdate.Month() == time.February
}

// GetName returns the formatted name of the contact.
func (model *Contact) GetName() string {
	var name string

	if model.Nickname != "" {
		name += model.Nickname

		if model.Firstname != "" {
			name += " (" + model.Firstname
		}

		if model.Lastname != "" {
			name += " " + model.Lastname + ")"
		} else {
			name += ")"
		}
	} else {
		name = model.Firstname

		if model.Lastname != "" {
			name += " " + model.Lastname
		}
	}

	return name
}
