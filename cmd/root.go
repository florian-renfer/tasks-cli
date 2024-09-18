package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var persistenceAdapter string

func init() {
	rootCmd.Flags().StringVarP(&persistenceAdapter, "source", "s", "csv", "Configure the persistence adapter to access the data source.")
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a minimal cli to manage your tasks.",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
