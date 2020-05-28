package birthday

import (
	"strconv"
	"strings"
	"time"

	"gobirthday/internal/models"

	"github.com/fallais/gonotify/pkg/notifiers"
	"github.com/sirupsen/logrus"
)

// DefaultHeader is the default header message.
const DefaultHeader = "Greets !"

// DefaultBase is the default base message.
const DefaultBase = "This is the birthay of [contact] !"

// DefaultAge is the default age message.
const DefaultAge = "[age] years old ! :)"

// DefaultFooter is the default footer message.
const DefaultFooter = "Bye !"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// GoBirthday is a birthday reminder that helps you to not forget your loved ones.
type GoBirthday struct {
	contacts             []*models.Contact
	notifiers            []notifiers.Notifier
	notificationTemplate map[string]string
	handleLeapYears      bool
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewGoBirthday returns new GoBirthday.
func NewGoBirthday(handleLeapYears bool, notificationTemplate map[string]string, contacts []*models.Contact, notifiers []notifiers.Notifier) *GoBirthday {
	gb := &GoBirthday{
		contacts:             contacts,
		notifiers:            notifiers,
		notificationTemplate: notificationTemplate,
		handleLeapYears:      handleLeapYears,
	}

	// Check the template
	if _, ok := notificationTemplate["header"]; !ok {
		notificationTemplate["header"] = DefaultHeader
	}
	if _, ok := notificationTemplate["base"]; !ok {
		notificationTemplate["base"] = DefaultBase
	}
	if _, ok := notificationTemplate["age"]; !ok {
		notificationTemplate["age"] = DefaultAge
	}
	if _, ok := notificationTemplate["footer"]; !ok {
		notificationTemplate["footer"] = DefaultFooter
	}

	return gb
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Notify notifies all the birthdays that need to be wished.
func (gb *GoBirthday) Notify() {
	logrus.WithFields(logrus.Fields{
		"nb_contacts": len(gb.contacts),
	}).Infoln("Checking all the contacts")

	// Check all the contacts
	wished := 0
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
				message := ""

				// Add header
				message += gb.notificationTemplate["header"]
				message += " "

				// Add base
				message += gb.notificationTemplate["base"]
				message += " "

				// Add age if not null
				if contact.GetAge() != 0 {
					message += gb.notificationTemplate["age"]
					message += " "
				}

				// Add footer
				message += gb.notificationTemplate["footer"]

				// Replace values
				r := strings.NewReplacer("{{contact}}", contact.Firstname, "{{age}}", strconv.Itoa(contact.GetAge()))
				message = r.Replace(message)

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

			wished++
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

	logrus.WithFields(logrus.Fields{
		"nb_contacts": len(gb.contacts),
		"nb_whished":  wished,
	}).Infoln("All the contacts have been checked")
}

// NbContacts return the number of contacts.
func (gb *GoBirthday) NbContacts() int {
	return len(gb.contacts)
}

// NbNotifiers return the number of notifiers.
func (gb *GoBirthday) NbNotifiers() int {
	return len(gb.notifiers)
}
