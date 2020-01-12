package cmd

import (
	"fmt"
	"os"

	"github.com/dnnrly/httpref"
	"github.com/spf13/cobra"
)

var all = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httpref [item]",
	Short: "Command line access to HTTP references",
	Long: `This displays useful information related to HTTP.

Most of the content comes from the Mozilla developer documentation (https://developer.mozilla.org/en-US/docs/Web/HTTP) and is copyright Mozilla and individual contributors. See https://developer.mozilla.org/en-US/docs/MDN/About#Copyrights_and_licenses for details.`,
	RunE: root,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&all, "all", "", all, "List all of the summaries available")
}

func root(cmd *cobra.Command, args []string) error {
	results := httpref.Statuses

	if !all {
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "Must specify a status code, or part of it\n")
			os.Exit(1)
		} else {
			results = results.ByName(args[0])
		}
	}

	switch len(results) {
	case 0:
		fmt.Fprintf(os.Stderr, "Name note recognised\n")
		os.Exit(1)
	case 1:
		fmt.Printf("%s - %s\n\n%s\n", results[0].Name, results[0].Summary, results[0].Description)
	default:
		for _, r := range results {
			fmt.Printf("\t%s\t%s\n", r.Name, r.Summary)
		}
	}

	return nil
}
