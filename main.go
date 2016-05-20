package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Sirupsen/logrus"
)

// Contact is a contact struct
type Contact struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
}

// ParseJSON contacts file
func ParseJSON(filename string) ([]Contact, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var contacts []Contact

	json.Unmarshal(file, &contacts)

	return contacts, nil
}

func main() {
	// Parse the contacts file
	list, err := ParseJSON("contacts.json")
	if err != nil {
		logrus.Fatal("Error while parsing the JSON file : ", err)
	}
	fmt.Printf("Results: %v\n", list)

	// Search for a birthday to wish
	for _, contact := range list {
		parsedBirthdate, err := time.Parse("02/01/2006", contact.Birthdate)
		if err != nil {
			logrus.Errorln("Error while parsing the JSON file : ", err)
		}

		if parsedBirthdate.Day() == time.Now().Day() && parsedBirthdate.Month() == time.Now().Month() {
			logrus.Infoln("Today it is", contact.Firstname, "birthday !")
		}
	}
}
