package models

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/sirupsen/logrus"
)

// BirthdateRegex is the Regex used to parse the birthdate.
const BirthdateRegex = "(\\d{2})\\/(\\d{2})\\/?(\\d{4})?"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Contact is a contact.
type Contact struct {
	Firstname   string
	Lastname    string
	Nickname    string
	Description string
	Birthdate   string

	day   int
	month int
	year  int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewContact returns a valid Contact instance
func NewContact(firstname, lastname, nickname, description, birthdate string) (*Contact, error) {
	// Create the model
	model := &Contact{
		Firstname:   firstname,
		Lastname:    lastname,
		Nickname:    nickname,
		Description: description,
		Birthdate:   birthdate,
	}

	// Validate the model
	err := model.Validate()
	if err != nil {
		return nil, err
	}

	// Sanitize the birthdate
	err = model.sanitizeBirthdate()
	if err != nil {
		return nil, err
	}

	return model, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Sanitize the birthdate
func (model *Contact) sanitizeBirthdate() error {
	// Compile the regex
	r, err := regexp.Compile(BirthdateRegex)
	if err != nil {
		return fmt.Errorf("error while compiling the regex : %s", err)
	}

	// Find the values
	subs := r.FindStringSubmatch(model.Birthdate)

	// Convert the day
	day, err := strconv.Atoi(subs[1])
	if err != nil {
		return fmt.Errorf("Error while converting the day : %s", err)
	}

	// Convert the month
	month, err := strconv.Atoi(subs[2])
	if err != nil {
		return fmt.Errorf("Error while converting the month : %s", err)
	}

	// Convert the year if it exists
	var year int
	if subs[3] != "" {
		year, err = strconv.Atoi(subs[3])
		if err != nil {
			logrus.Errorf("Error while converting the year : %s", err)
		}
	} else {
		year = 0
	}

	model.day = day
	model.month = month
	model.year = year

	return nil
}

// Validate the model.
func (model *Contact) Validate() error {
	return validation.ValidateStruct(model,
		validation.Field(&model.Firstname, validation.Required),
		validation.Field(&model.Birthdate, validation.Required, validation.Match(regexp.MustCompile(BirthdateRegex))),
	)
}

// GetAge returns the age of the contact. 0 if there is no Year.
func (model *Contact) GetAge() int {
	if model.year == 0 {
		return 0
	}

	return time.Now().Year() - model.year
}

// IsBirthdayToday returns true if it is the birthday of the contact.
func (model *Contact) IsBirthdayToday() bool {
	if model.day == time.Now().Day() && model.month == int(time.Now().Month()) {
		return true
	}

	return false
}

// IsBornOnLeapYear returns true if the contact is born on Frebruary the 29th.
func (model *Contact) IsBornOnLeapYear() bool {
	return model.day == 29 && model.month == int(time.February)
}
