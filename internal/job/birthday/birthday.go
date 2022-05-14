package birthday

import (
	"bytes"
	"fmt"
	"html/template"

	"hermes/internal/models"
	"hermes/pkg/notifiers"

	"github.com/sirupsen/logrus"
)

// DefaultHour is the default hour for wishing the birthday.
const DefaultHour = 10

// DefaultHour is the default minute for wishing the birthday.
const DefaultMinute = 30

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Birthday is the birthday to wish.
type Birthday struct {
	contact          *models.Contact
	notifiers        []notifiers.Notifier
	leapYearsEnabled bool
	leapYearsMode    string
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
	return fmt.Sprintf("%d %d %d %d *", DefaultMinute, DefaultHour, gb.contact.Birthdate.Day(), gb.contact.Birthdate.Month())
}

// Run is the convenient function for notify.
func (gb *Birthday) Run() {
	// Parse the template
	tmpl, err := template.New("birthday").Parse(DefaultTemplate)
	if err != nil {
		logrus.WithError(err).Errorln("error while parsing template")
		return
	}

	// Create the buffer
	buf := &bytes.Buffer{}

	// Prepare the data
	data := TemplateData{
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
		err := notifier.Notify(buf.String())
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
