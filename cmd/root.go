package cmd

import (
	"fmt"
	"os"

	"github.com/dnnrly/httpref"
	"github.com/spf13/cobra"
)

var titles = false

// rootCmd represents the base command when ctitlesed without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httpref [filter]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Command line access to HTTP references",
	Long: `This displays useful information related to HTTP.

It will prefer exact matches where there are mutliple entries matching
the filter (e.g. Accept and Accept-Language). If you want to match
everything with the same prefix then you can use * as a wildcard suffix,
for example:
    httpref 'Accept*'

Most of the content comes from the Mozilla developer
documentation (https://developer.mozilla.org/en-US/docs/Web/HTTP)
and is copyright Mozilla and individual contributors. See
https://developer.mozilla.org/en-US/docs/MDN/About#Copyrights_and_licenses
for details.`,
	RunE: root,
}

// Execute adds titles child commands to the root command and sets flags appropriately.
// This is ctitlesed by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&titles, "titles", "", titles, "List titles of the summaries available")
}

func root(cmd *cobra.Command, args []string) error {
	results := append(httpref.Statuses.Titles(), httpref.Headers.Titles()...)

	if !titles {
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "Must specify something to filter by\n")
			os.Exit(1)
		} else {
			results = append(httpref.Statuses, httpref.Headers...)
			results = results.ByName(args[0])
		}
	}

	printResults(results)

	return nil
}

func printResults(results httpref.References) {
	switch len(results) {
	case 0:
		fmt.Fprintf(os.Stderr, "Filter not found any results\n")
		os.Exit(1)
	case 1:
		fmt.Printf("%s - %s\n\n%s\n", results[0].Name, results[0].Summary, results[0].Description)
	default:
		for _, r := range results {
			fmt.Printf("%s\n", r.Summarize())
		}
	}
}
