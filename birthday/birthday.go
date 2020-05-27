package birthday

import (
	"time"

	"gobirthday/models"
	"gobirthday/providers"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// GoBirthday is a birthday reminder that allows you to not forget your loved ones.
type GoBirthday struct {
	contacts        []*models.Contact
	providers       []providers.Provider
	handleLeapYears bool
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewGoBirthday returns new GoBirthday.
func NewGoBirthday(handleLeapYears bool, contacts []*models.Contact, providers []providers.Provider) *GoBirthday {
	return &GoBirthday{
		contacts:        contacts,
		providers:       providers,
		handleLeapYears: handleLeapYears,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Notify notifies all the birthdays that need to be wished.
func (gb *GoBirthday) Notify() {
	// Check all the contacts
	logrus.Infoln("Check all the contacts")
	for _, contact := range gb.contacts {
		// Check the birthdate
		if contact.IsBirthdayToday() {
			logrus.WithFields(logrus.Fields{
				"age":       contact.GetAge(),
				"firstname": contact.Firstname,
				"lastname":  contact.Lastname,
			}).Infoln("Birthday to wish !")

			// Send all the notifications
			for _, provider := range gb.providers {
				logrus.WithFields(logrus.Fields{
					"provider_type":   provider.Type(),
					"provider_vendor": provider.Vendor(),
				}).Infoln("Sending the notification")
				err := provider.SendNotification(contact)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"provider_type":   provider.Type(),
						"provider_vendor": provider.Vendor(),
					}).Errorln("Error while sending the notification :", err)
					continue
				}

				logrus.WithFields(logrus.Fields{
					"provider_type":   provider.Type(),
					"provider_vendor": provider.Vendor(),
				}).Infoln("Successfully sent the notification")
			}
		}

		// Check leap years
		if gb.handleLeapYears && contact.IsBornOnLeapYear() && time.Now().Day() == 1 && time.Now().Month() == time.March {
			logrus.WithFields(logrus.Fields{
				"age":       contact.GetAge(),
				"firstname": contact.Firstname,
				"lastname":  contact.Lastname,
			}).Infoln("Birthday to wish on a leap year !")
		}
	}
	logrus.Debugln("All the contacts have been checked")
}

// NbContacts return the number of contacts.
func (gb *GoBirthday) NbContacts() int {
	return len(gb.contacts)
}

// NbProviders return the number of providers.
func (gb *GoBirthday) NbProviders() int {
	return len(gb.providers)
}
