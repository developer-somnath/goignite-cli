package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const GoIgniteVersion = "1.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GoIgnite",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("GoIgnite version %s\n", GoIgniteVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
