package birthday

import (
	"time"
	"os"
	"os/signal"

	"gobirthday/providers"
	"gobirthday/models"

	"github.com/sirupsen/logrus"
	"github.com/robfig/cron"
)

// BirthdateDefaultFormat is the birthdate format.
const BirthdateDefaultFormat = "02/01/2006"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// GoBirthday is a birthday reminder that allows you to not forget your loved ones.
type GoBirthday struct {
	contacts         []*models.Contact
	providers        []providers.Provider
	cronExp string
	cron *cron.Cron
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewGoBirthday returns new GoBirthday with the given CRON expression.
func NewGoBirthday(cronExp string) (*GoBirthday, error) {
	// Create the object
	gb := &GoBirthday{
		cron: cron.New(),
		cronExp: cronExp,
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
		// Check the birthdate
		if contact.Birthdate.Day == time.Now().Day() && contact.Birthdate.Month == int(time.Now().Month()) {
			// Calculate the age
			age := calculateAge(contact.Birthdate)

			logrus.WithFields(logrus.Fields{
				"age":   age,
				"firstname": contact.Firstname,
				"lastname": contact.Lastname,
			}).Infoln("Birthday to wish !")

			// Send all the notifications
			for _, provider := range gb.providers {
				err := provider.SendNotification(contact.Firstname, contact.Lastname, age)
				if err != nil {
					logrus.Errorln("Error while sending the notification :", err)
					continue
				}
			}
		}
	}
}

// NbContacts return the number of contacts.
func (gb *GoBirthday) NbContacts() int {
	return len(gb.contacts)
}

// NbProviders return the number of providers.
func (gb *GoBirthday) NbProviders() int {
	return len(gb.providers)
}

// Start starts the program.
func (gb *GoBirthday) Start() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)

	// Add the function to the CRON
	gb.Notify()
	logrus.WithFields(logrus.Fields{
		"cron_exp":   gb.cronExp,
	}).Infoln("Adding function to the CRON")
	gb.cron.AddFunc(gb.cronExp, gb.Notify)

	// Start the CRON
	logrus.Infoln("Starting the CRON")
	gb.cron.Start()

	// Handle KILL or CTRL+C
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	go func() {
		for range signalChan {
			logrus.Infoln("Received an interrupt, stopping services...")

			gb.cron.Stop()

			logrus.Infoln("Services stopped")

			cleanupDone <- true
		}
	}()

	logrus.Infoln("Waiting for birthdays to wish")

	<-cleanupDone
}