package main

import (
	"flag"
	"gobirthday/birthday"
	"time"
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/robfig/cron"
)

var (
	logging       = flag.String("logging", "info", "Logging level")
	contactsFile  = flag.String("contacts_file", "contacts.json", "Contacts")
	providersFile = flag.String("providers_file", "providers.json", "Providers")
	cronExp = flag.String("cron_exp", "0 8 * * *", "Cron ?")
	signalChan   = make(chan os.Signal, 1)
	cleanupDone  = make(chan bool)
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
	gb, err := birthday.NewGoBirthday(*contactsFile, *providersFile, birthday.BirthdateDefaultFormat)
	if err != nil {
		logrus.Fatalln("Error while creating the GoBirthday : ", err)
	}

	// Create the CRON
	logrus.Infoln("Creating the CRON")
	c := cron.New()
	c.AddFunc(*cronExp, gb.Notify)
	c.Start()
	logrus.Infoln("Successfully created the CRON")

	// Handle KILL or CTRL+C
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	go func() {
		for range signalChan {
			logrus.WithFields(logrus.Fields{
				"channel": "system",
			}).Infoln("Received an interrupt, stopping services...")

			c.Stop()

			logrus.WithFields(logrus.Fields{
				"channel": "system",
			}).Infoln("Services stopped")

			cleanupDone <- true
		}
	}()

	<-cleanupDone
}
