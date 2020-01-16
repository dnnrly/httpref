package cmd

import (
	"github.com/dnnrly/httpref"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status [filter]",
	Short: "References for HTTP status codes",
	Run: func(cmd *cobra.Command, args []string) {
		results := httpref.Statuses

		if len(args) == 0 {
			results = results.Titles()
		} else {
			results = results.ByName(args[0])
		}

		printResults(results)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
