package main

import (
	"os"
	"os/signal"
	"flag"
	"gobirthday/birthday"

	"github.com/sirupsen/logrus"
	"github.com/robfig/cron"
)

var (
	logging         = flag.String("logging", "info", "Logging level")
	contactsFile    = flag.String("contacts_file", "contacts.json", "Contacts list")
	providersFile   = flag.String("providers_file", "providers.json", "Providers list")
	handleLeapYears = flag.Bool("handle_leap_years", false, "Handle leap years ?")
	cronExp         = flag.String("cron_exp", "0 30 14 * * *", "Cron ?")
	runOnStartup    = flag.Bool("run_on_startup", false, "Run on startup ?")
)

func init() {
	// Parse the flags
	flag.Parse()

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
	logrus.WithFields(logrus.Fields{
		"cron_exp": *cronExp,
		"handle_leap_years": *handleLeapYears,
		"run_on_startup": *runOnStartup,
	}).Infoln("Creating the GoBirthday")
	gb := birthday.NewGoBirthday(*handleLeapYears)
	logrus.Infoln("Successfully created the GoBirthday")

	// Add the contacts
	logrus.Infoln("Adding the contacts")
	err := gb.AddContacts(*contactsFile)
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

	// Create the channels
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)

	// Run on startup
	if *runOnStartup {
		gb.Notify()
	}

	c := cron.New()

	// Add the function to the CRON
	logrus.WithFields(logrus.Fields{
		"cron_exp": *cronExp,
	}).Infoln("Adding function to the CRON")
	c.AddFunc(*cronExp, gb.Notify)

	// Handle KILL or CTRL+C
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	go func() {
		for _ = range signalChan {
			logrus.Infoln("Received an interrupt, stopping...")
			cleanupDone <- true
		}
	}()

	logrus.Infoln("Waiting for birthdays to wish")

	<-cleanupDone
}
