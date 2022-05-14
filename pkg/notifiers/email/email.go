package email

import (
	"fmt"

	"hermes/pkg/notifiers"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Type is the type of the provider.
const Type = "Email"

// Vendor is the vendor of the provider.
const Vendor = "Email"

type email struct {
	host      string
	port      int
	recipient string
	subject   string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewNotifier returns a new notifier for SMTP.
func NewNotifier(settings map[string]interface{}) notifiers.Notifier {
	// Initial values
	host := ""
	port := 25
	recipient := ""
	subject := ""

	// Process the values
	for key, value := range settings {
		switch key {
		case "host":
			host = value.(string)
		case "port":
			port = int(value.(float64))
		case "recipient":
			recipient = value.(string)
		case "subject":
			subject = value.(string)
		default:
			logrus.WithFields(logrus.Fields{
				"key":   key,
				"value": value,
			}).Infoln("Wrong setting for SMTP")
		}
	}

	return &email{
		host:      host,
		port:      port,
		recipient: recipient,
		subject:   subject,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Notify sends a notification.
func (n *email) Notify(msg string) error {
	// Create the message
	m := gomail.NewMessage()
	m.SetHeader("From", n.recipient)
	m.SetHeader("To", n.recipient)
	m.SetHeader("Subject", n.subject)
	m.SetBody("text/html", msg)

	// Dial
	d := gomail.Dialer{Host: n.host, Port: n.port}

	// Send the email
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}

// Name returns the formatted name.
func (n *email) Name() string {
	return fmt.Sprintf("%s-%s", Type, Vendor)
}
