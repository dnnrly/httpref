package cmd

import (
	"github.com/dnnrly/httpref"
	"github.com/spf13/cobra"
)

// methodsCmd represents the headers command
var methodsCmd = &cobra.Command{
	Use:   "methods",
	Short: "References for common HTTP methods",
	Run: func(cmd *cobra.Command, args []string) {
		results := httpref.Methods

		if len(args) == 0 {
			results = results.Titles()
		} else {
			results = results.ByName(args[0])
		}

		printResults(results)
	},
}

func init() {
	rootCmd.AddCommand(methodsCmd)
}
