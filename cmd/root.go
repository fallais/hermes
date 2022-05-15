package cmd

import (
	"os"

	"hermes/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "Hermes is a tool written in Go that reminds you things you have to do and birthday !",
	Long:  ``,
	Run:   internal.Run,
}

func init() {
	rootCmd.Flags().StringP("config", "c", "/config.yaml", "Configuration file")
}

// Execute the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
