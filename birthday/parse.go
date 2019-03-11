package birthday

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"

	"gobirthday/models"
	"gobirthday/providers/email"
	"gobirthday/providers/sms/free"
	"gobirthday/providers/sms/orange"

	"github.com/sirupsen/logrus"
)

// Contact is a contact.
type Contact struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Birthdate   string `json:"birthdate"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
}

// Provider is a provider.
type Provider struct {
	Type     string                 `json:"type"`
	Vendor   string                 `json:"vendor"`
	Settings map[string]interface{} `json:"settings"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// AddContacts adds the contacts.
func (gb *GoBirthday) AddContacts(filename string) error {
	var contacts []*Contact

	// Read the configuration file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error while reading the file : %s", err)
	}

	// Unmarshal the configuration file
	err = json.Unmarshal(file, &contacts)
	if err != nil {
		return fmt.Errorf("error while unmarshalling : %s", err)
	}

	// Process the contacts
	for _, contact := range contacts {
		// Compile the regex
		r, err := regexp.Compile("(\\d{2})\\/(\\d{2})\\/?(\\d{4})?")
		if err != nil {
			return fmt.Errorf("error while compiling the regex : %s", err)
		}

		// Check the birthdate
		if !r.MatchString(contact.Birthdate) {
			logrus.WithFields(logrus.Fields{
				"firstname": contact.Firstname,
				"lastname":  contact.Lastname,
				"birthdate": contact.Birthdate,
			}).Errorln("The birthdate is incorrect")

			continue
		}

		// Find the values
		subs := r.FindStringSubmatch(contact.Birthdate)

		// Convert the day
		day, err := strconv.Atoi(subs[1])
		if err != nil {
			logrus.Errorf("Error while converting the day : %s", err)
			continue
		}

		// Convert the month
		month, err := strconv.Atoi(subs[2])
		if err != nil {
			logrus.Errorf("Error while converting the month : %s", err)
			continue
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

		// Create the birthdate
		bd := models.NewBirthdate(day, month, year)

		// Create the contact
		c := models.NewContact(contact.Firstname, contact.Lastname, contact.Nickname, contact.Description, bd)

		// Add the contact
		gb.contacts = append(gb.contacts, c)
	}

	return nil
}

// AddProviders ...
func (gb *GoBirthday) AddProviders(filename string) error {
	var providers []*Provider

	// Read the configuration file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error while reading the providers file : %s", err)
	}

	// Unmarshal the configuration file
	err = json.Unmarshal(file, &providers)
	if err != nil {
		return fmt.Errorf("Error while unmarshalling the providers : %s", err)
	}

	// Create the providers
	for _, provider := range providers {
		switch provider.Type {
		case "sms":
			switch provider.Vendor {
			case "free":
				freeProvider := free.NewProvider(provider.Settings)
				gb.providers = append(gb.providers, freeProvider)
				break
			case "orange":
				orangeProvider := orange.NewProvider(provider.Settings)
				gb.providers = append(gb.providers, orangeProvider)
			default:
				return fmt.Errorf("Wrong vendor of SMS provider : %s", provider.Type)
			}
		case "email":
			emailProvider := email.NewProvider(provider.Settings)
			gb.providers = append(gb.providers, emailProvider)
			break
		default:
			return fmt.Errorf("Wrong type of provider : %s", provider.Type)
		}
	}

	return nil
}
