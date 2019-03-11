package main

import (
	"flag"
	"gobirthday/birthday"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	logging         = flag.String("logging", "info", "Logging level")
	contactsFile    = flag.String("contacts_file", "contacts.json", "Contacts list")
	providersFile   = flag.String("providers_file", "providers.json", "Providers list")
	handleLeapYears = flag.Bool("handle_leap_years", false, "Handle leap years ?")
	cronExp         = flag.String("cron_exp", "50 15 * * *", "Cron ?")
	runOnStartup = flag.Bool("run_on_startup", false, "Run on startup ?")
)

func init() {
	// Parse the flags
	flag.Parse()

	// Set localtime to UTC
	time.Local = time.UTC

	// Set the logging level
	switch *logging {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Set the TextFormatter
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	logrus.Infoln("gobirthday is starting")
}

func main() {
	// Parse the contacts file
	logrus.Infoln("Creating the GoBirthday")
	gb, err := birthday.NewGoBirthday(*cronExp, *handleLeapYears, *runOnStartup)
	if err != nil {
		logrus.Fatalf("Error while creating the GoBirthday : %s", err)
	}

	// Add the contacts
	logrus.Infoln("Adding the contacts")
	err = gb.AddContacts(*contactsFile)
	if err != nil {
		logrus.Fatalf("Error while adding the contacts : %s", err)
	}
	logrus.WithFields(logrus.Fields{
		"nb_contacts": gb.NbContacts(),
	}).Infoln("Successfully added the contacts")

	// Add the providers
	logrus.Infoln("Adding the providers")
	err = gb.AddProviders(*providersFile)
	if err != nil {
		logrus.Fatalf("Error while adding the providers : %s", err)
	}
	logrus.WithFields(logrus.Fields{
		"nb_providers": gb.NbProviders(),
	}).Infoln("Successfully added the providers")

	// Start
	gb.Start()
}
