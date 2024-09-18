package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Initialize cobra command hierarchy
func init() {
	rootCmd.AddCommand(versionCmd)
}

// `version` command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tasks-cli.",
	Long:  `All software has versions. This is task-cli's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tasks-cli version 0.0.1")
	},
}
