package internal

import (
	"fmt"

	"gobirthday/internal/models"

	"github.com/fallais/gonotify/pkg/notifiers"
	"github.com/fallais/gonotify/pkg/notifiers/email"
	"github.com/fallais/gonotify/pkg/notifiers/sms/free"
	"github.com/fallais/gonotify/pkg/notifiers/sms/orange"
	"github.com/spf13/viper"
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

func setupContacts() ([]*models.Contact, error) {
	var contacts []*models.Contact
	var configContacts []*Contact

	err := viper.UnmarshalKey("contacts", &configContacts)
	if err != nil {
		return nil, err
	}

	// Process the contacts
	for _, configContact := range configContacts {
		// Create the contact
		c, err := models.NewContact(configContact.Firstname, configContact.Lastname, configContact.Nickname, configContact.Description, configContact.Birthdate)
		if err != nil {
			return nil, err
		}

		// Add the contact
		contacts = append(contacts, c)
	}

	return contacts, nil
}

func setupProviders() ([]notifiers.Notifier, error) {
	var providers []notifiers.Notifier
	var configProviders []*Provider

	err := viper.UnmarshalKey("providers", &configProviders)
	if err != nil {
		return nil, err
	}

	// Create the providers
	for _, configProvider := range configProviders {
		switch configProvider.Type {
		case "sms":
			switch configProvider.Vendor {
			case "free":
				freeProvider := free.NewNotifier(configProvider.Settings)
				providers = append(providers, freeProvider)
				break
			case "orange":
				orangeProvider := orange.NewNotifier(configProvider.Settings)
				providers = append(providers, orangeProvider)
			default:
				return nil, fmt.Errorf("Wrong vendor of SMS provider : %s", configProvider.Vendor)
			}
		case "email":
			emailProvider := email.NewNotifier(configProvider.Settings)
			providers = append(providers, emailProvider)
			break
		default:
			return nil, fmt.Errorf("Wrong type of provider : %s", configProvider.Type)
		}
	}

	return providers, nil
}
