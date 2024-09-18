package cmd

import (
	"github.com/florian-renfer/tasks/internal/task"
	"github.com/spf13/cobra"
)

// Label used for task generation
var label string

// Initialize cobra command hierarchy
func init() {
	rootCmd.AddCommand(addCmd)
}

// `add` command
var addCmd = &cobra.Command{
	Use:       "add [label]",
	Short:     "Add a new task to the list of tasks managed by tasks-cli.",
	Long:      `Create, store and add a new task to the list of tasks managed by tasks-cli. The first argument is used as 'label' for the task created. In case 'label' contains spaces remember to escape them by wrapping the entire value in double quoutes.`,
	ValidArgs: []string{"label"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		s := task.GetCsvTaskServiceInstance()
		s.Create(label)
	},
}
