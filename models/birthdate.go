package models

import (
	"fmt"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Birthdate is the structure of a birthdate.
type Birthdate struct {
	Day   int
	Month int
	Year  int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewBirthdate returns a new Birthdate.
func NewBirthdate(day, month, year int) *Birthdate {
	return &Birthdate{
		Day:   day,
		Month: month,
		Year:  year,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ToString returns a string representation of the birthdate.
func (m *Birthdate) ToString() string {
	if m.Year == 0 {
		return fmt.Sprintf("%02d/%02d", m.Day, m.Month)
	}

	return fmt.Sprintf("%02d/%02d/%02d", m.Day, m.Month, m.Year)
}
