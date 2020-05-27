package birthday

import (
	"time"

	"gobirthday/internal/models"

	"github.com/fallais/gonotify/pkg/notifiers"
	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// GoBirthday is a birthday reminder that helps you to not forget your loved ones.
type GoBirthday struct {
	contacts        []*models.Contact
	notifiers       []notifiers.Notifier
	handleLeapYears bool
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewGoBirthday returns new GoBirthday.
func NewGoBirthday(handleLeapYears bool, contacts []*models.Contact, notifiers []notifiers.Notifier) *GoBirthday {
	return &GoBirthday{
		contacts:        contacts,
		notifiers:       notifiers,
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
			for _, notifier := range gb.notifiers {
				message := "Coucou"

				logrus.WithFields(logrus.Fields{
					"notifier": notifier.Name(),
				}).Infoln("Sending the notification")
				err := notifier.Notify(message)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"notifier": notifier.Name(),
					}).Errorln("Error while sending the notification :", err)
					continue
				}

				logrus.WithFields(logrus.Fields{
					"notifier": notifier.Name(),
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

// NbNotifiers return the number of notifiers.
func (gb *GoBirthday) NbNotifiers() int {
	return len(gb.notifiers)
}
