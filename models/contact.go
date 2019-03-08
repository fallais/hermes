package models

// Birthdate is the structure of a birthdate.
type Birthdate struct {
	Day int
	Month int
	Year int
}

// Contact is a contact.
type Contact struct {
	Firstname string
	Lastname  string
	Description string
	Birthdate Birthdate
}