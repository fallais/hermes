package birthday

import (
	"bytes"
	"fmt"
	"html/template"

	"hermes/internal/models"

	"github.com/fallais/gonotify/pkg/notifiers"
	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type Birthday struct {
	contact          *models.Contact
	notifiers        []notifiers.Notifier
	leapYearsEnabled bool
	leapYearsMode    string
}

type data struct {
	contact string
	age     int
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New returns new GoBirthday.
func New(leapYearsEnabled bool, leapYearsMode string, contact *models.Contact, notifiers []notifiers.Notifier) *Birthday {
	gb := &Birthday{
		contact:          contact,
		notifiers:        notifiers,
		leapYearsEnabled: leapYearsEnabled,
		leapYearsMode:    leapYearsMode,
	}

	return gb
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetCRONExpression returns the calculated CRON expression for the birthday.
func (gb *Birthday) GetCRONExpression() string {
	return fmt.Sprintf("0 0 %d %d *", gb.contact.Birthdate.Day(), gb.contact.Birthdate.Month())
}

// Run is the convenient function for notify.
func (gb *Birthday) Run() {
	// Parse the template
	tmpl, err := template.New("birthday").Parse(MessageTemplate)
	if err != nil {
		logrus.WithError(err).Errorln("error while parsing template")
		return
	}

	// Create the buffer
	buf := &bytes.Buffer{}

	// Prepare the data
	data := data{
		contact: gb.contact.GetName(),
		age:     gb.contact.GetAge(),
	}

	// Execute the template with data
	err = tmpl.Execute(buf, data)
	if err != nil {
		logrus.WithError(err).Errorln("error while executing template")
		return
	}

	// Send all the notifications
	for _, notifier := range gb.notifiers {
		logrus.WithFields(logrus.Fields{
			"notifier": notifier.Name(),
		}).Infoln("Sending the notification")
		err := notifier.Notify(fmt.Sprintf("Do not forget your friend %s born on a leap year !", gb.contact.Firstname))
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
