package birthday

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gobirthday/providers/email"
	"gobirthday/providers/sms/free"
	"gobirthday/providers/sms/orange"
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
