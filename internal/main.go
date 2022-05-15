package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"hermes/internal/reminder/birthday"
	"hermes/internal/reminder/thing"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	// Flag
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		logrus.WithError(err).Fatalln("Error with the configuration file flag")
	}

	// Read configuration file
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.WithError(err).Fatalln("Error when reading configuration file")
	}

	// Initialize values with Viper
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		logrus.WithError(err).Fatalln("Error when reading configuration data")
	}

	// Setup the contacts
	contacts, err := setupContacts()
	if err != nil {
		logrus.WithError(err).Fatalln("Error when setup the contacts")
	}

	// Setup the things
	things, err := setupThings()
	if err != nil {
		logrus.WithError(err).Fatalln("Error when setup the things")
	}

	// Setup the providers
	providers, err := setupProviders()
	if err != nil {
		logrus.WithError(err).Fatalln("Error when setup the providers")
	}
	logrus.WithFields(logrus.Fields{
		"nb_providers": len(providers),
	}).Infoln("Successfully setup the providers")

	// Create the CRON
	c := cron.New()

	// Add birthdays to the CRON
	logrus.WithFields(logrus.Fields{
		"nb_contacts": len(contacts),
	}).Infoln("Adding birthdays to the CRON")
	for _, contact := range contacts {
		// Create the birthday
		b := birthday.New(false, "", *contact, providers)

		// Add the birthday
		logrus.WithFields(logrus.Fields{
			"contact": contact.GetName(),
		}).Infoln("Adding birthday to the CRON")
		_, err := c.AddJob(b.GetCRONExpression(), b)
		if err != nil {
			logrus.WithError(err).Fatalln("error while adding the birthday to the CRON")
		}
	}

	// Add things to the CRON
	logrus.WithFields(logrus.Fields{
		"nb_things": len(things),
	}).Infoln("Adding things to the CRON")
	for _, t := range things {
		th := thing.New(t.Name, t.When, providers)

		logrus.WithFields(logrus.Fields{
			"name": t.Name,
			"when": t.When,
		}).Infoln("Adding the thing to the CRON")
		_, err := c.AddJob(th.GetCRONExpression(), th)
		if err != nil {
			logrus.WithError(err).Fatalln("error while adding the thing to the CRON")
		}
	}

	// Start the CRON
	logrus.Infoln("Starting the CRON")
	c.Start()
	logrus.WithFields(logrus.Fields{
		"nb_entries": len(c.Entries()),
	}).Infoln("CRON has been started")

	// Serve
	logrus.WithFields(logrus.Fields{
		"port": ":8000",
	}).Infoln("Starting the Web server")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		logrus.WithError(err).Fatalln("Error while starting the Web server")
	}
}
