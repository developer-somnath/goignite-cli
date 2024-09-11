package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goignite",
	Short: "GoIgnite is a CLI tool for managing your framework",
	Long:  "A framework generator to create new projects, controllers, models, and more.",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
