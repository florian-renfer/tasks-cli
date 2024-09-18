package cmd

import (
	"github.com/florian-renfer/tasks/internal/task"
	"github.com/spf13/cobra"
)

var all bool

// Initialize `list` command
func init() {
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "List all tasks without regard to their state.")
	rootCmd.AddCommand(listCmd)
}

// `list` command configuration for cobra
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks.",
	Long:  `Lists all tasks stored inside the local storage. By using proper flags the result can be modified.`,
	Run: func(cmd *cobra.Command, args []string) {
		task.List(all)
	},
}
