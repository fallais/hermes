package cmd

import (
	"fmt"
	"os"

	"hermes/internal"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "hermes",
	Short:             "Hermes is a tool written in Go that reminds you things you have to do and birthday !",
	Long:              ``,
	PersistentPreRunE: persistentPreRunE,
	Run:               internal.Run,
}

func persistentPreRunE(cmd *cobra.Command, args []string) error {
	logging, err := cmd.Flags().GetString("logging")
	if err != nil {
		return fmt.Errorf("error with the logging flag: %s", err)
	}

	// Parse the logging level
	level, err := logrus.ParseLevel(logging)
	if err != nil {
		return fmt.Errorf("error while parsing the logging level: %s", err)
	}

	// Set
	logrus.SetLevel(level)

	// Set the TextFormatter
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	return nil
}

func init() {
	rootCmd.PersistentFlags().StringP("logging", "l", "info", "Logging level")
	rootCmd.Flags().StringP("config", "c", "/config.yaml", "Configuration file")
}

// Execute the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
