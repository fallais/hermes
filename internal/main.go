package internal

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/signal"

	"hermes/internal/birthday"

	"github.com/robfig/cron"
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
	logrus.WithFields(logrus.Fields{
		"nb_contacts": len(contacts),
	}).Infoln("Successfully setup the contacts")

	// Setup the tasks
	// TODO

	// Setup the providers
	providers, err := setupProviders()
	if err != nil {
		logrus.WithError(err).Fatalln("Error when setup the providers")
	}
	logrus.WithFields(logrus.Fields{
		"nb_providers": len(providers),
	}).Infoln("Successfully setup the providers")

	// Create the channels
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)

	// Create the CRON
	c := cron.New()

	// Add birthdays to the CRON
	logrus.WithFields(logrus.Fields{
		"cron_exp": viper.GetString("general.cron_exp"),
	}).Infoln("Adding birthdays to the CRON")
	for _, contact := range contacts {
		birthday := birthday.New(false, "", contact, providers)
		c.AddFunc(birthday.GetCRONExpression(), birthday.Run)
	}

	// Start the CRON
	logrus.Infoln("Starting the CRON")
	c.Start()

	// Handle the kill signals
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	go func() {
		for range signalChan {
			logrus.Infoln("Received an interrupt, stopping...")

			// Stop the CRON
			c.Stop()

			cleanupDone <- true
		}
	}()

	logrus.Infoln("Waiting for birthdays to wish")

	<-cleanupDone
}
