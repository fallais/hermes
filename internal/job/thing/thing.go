package thing

import (
	"bytes"
	"html/template"

	"github.com/fallais/gonotify/pkg/notifiers"
	"github.com/sirupsen/logrus"
)

// DefaultHour is the default hour for wishing the birthday.
const DefaultHour = 10

// DefaultHour is the default minute for wishing the birthday.
const DefaultMinute = 30

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Thing is something that has to be done.
type Thing struct {
	name      string
	when      string
	notifiers []notifiers.Notifier
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New returns new GoBirthday.
func New(name, when string, notifiers []notifiers.Notifier) *Thing {
	gb := &Thing{
		name:      name,
		when:      when,
		notifiers: notifiers,
	}

	return gb
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetCRONExpression returns the calculated CRON expression for the birthday.
func (job *Thing) GetCRONExpression() string {
	return job.when
}

// Run is the convenient function for notify.
func (job *Thing) Run() {
	// Parse the template
	tmpl, err := template.New("thing").Parse(MessageTemplate)
	if err != nil {
		logrus.WithError(err).Errorln("error while parsing template")
		return
	}

	// Create the buffer
	buf := &bytes.Buffer{}

	// Prepare the data
	data := TemplateData{
		Thing: job.name,
	}

	// Execute the template with data
	err = tmpl.Execute(buf, data)
	if err != nil {
		logrus.WithError(err).Errorln("error while executing template")
		return
	}

	// Send all the notifications
	for _, notifier := range job.notifiers {
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
