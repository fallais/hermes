package birthday

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"gobirthday/providers"
	"gobirthday/providers/sms/free"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// GoBirthday is a birthday reminder.
type GoBirthday struct {
	contacts         []*Contact
	providers        []providers.Provider
	birthdateFormat  string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewGoBirthday returns a valid GoBirthday instance
func NewGoBirthday(contactsFile, providersFile, birthdateFormat string) (*GoBirthday, error) {
	var contacts []*Contact
	var providersTmp []Provider

	// Read the configuration file
	file, err := ioutil.ReadFile(contactsFile)
	if err != nil {
		return nil, fmt.Errorf("Error while reading the configuration file : %s", err)
	}

	// Unmarshal the configuration file
	err = json.Unmarshal(file, &contacts)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the configuration : %s", err)
	}

	// Read the configuration file
	file, err = ioutil.ReadFile(providersFile)
	if err != nil {
		return nil, fmt.Errorf("Error while reading the providers file : %s", err)
	}

	// Unmarshal the configuration file
	err = json.Unmarshal(file, &providersTmp)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the providers : %s", err)
	}

	// Create the providers
	var prvds []providers.Provider
	for _, provider := range providersTmp {
		switch provider.Type {
		case "sms":
			switch provider.Vendor {
			case "free":
				prvd := free.NewProvider(provider.Settings)
				prvds = append(prvds, prvd)
			default:
				return nil, fmt.Errorf("Wrong vendor of SMS provider : %s", provider.Type)
			}
		default:
			return nil, fmt.Errorf("Wrong type of provider : %s", provider.Type)
		}
	}

	// Create the object
	gb := &GoBirthday{
		contacts:  contacts,
		providers: prvds,
	}

	// Check the birthdate format
	if len(strings.TrimSpace(birthdateFormat)) == 0 {
		gb.birthdateFormat = BirthdateDefaultFormat
	} else {
		gb.birthdateFormat = birthdateFormat
	}

	return gb, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Notify notifies all the birthdays that need to be wished.
func (gb *GoBirthday) Notify() {
	// Process all the contacts
	for _, contact := range gb.contacts {
		// Parse the birthdate
		birthdateParsed, err := time.Parse(gb.birthdateFormat, contact.Birthdate)
		if err != nil {
			logrus.Errorln("Error while parsing the birthdate :", err)
			continue
		}

		// Check the birthdate
		if birthdateParsed.Day() == time.Now().Day() && birthdateParsed.Month() == time.Now().Month() {
			fmt.Println("Today it is", contact.Firstname, "birthday !", time.Now().Year()-birthdateParsed.Year(), "years old !")

			// Send all the notifications
			for _, provider := range gb.providers {
				err := provider.SendNotification(contact.Firstname, contact.Lastname, birthdateParsed)
				if err != nil {
					logrus.Errorln("Error while sending the notification :", err)
					continue
				}
			}
		}

	}
}
