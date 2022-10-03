package cmd

import (
	"fmt"
	"os"

	"github.com/dnnrly/httpref"
	"github.com/spf13/cobra"

	"github.com/dnnrly/paragraphical"
)

var (
	titles     = false
	width      = 100
	searchTerm = ""
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httpref [filter]",
	Args:  cobra.MaximumNArgs(1),
	Short: "Command line access to HTTP references",
	Long: paragraphical.Format(width, `This displays useful information related to HTTP.

It will prefer exact matches where there are multiple entries matching the filter (e.g. Accept and Accept-Language). If you want to match everything with the same prefix then you can use * as a wildcard suffix, for example:
    httpref 'Accept*'

Most of the content comes from the Mozilla developer documentation (https://developer.mozilla.org/en-US/docs/Web/HTTP) and is copyright Mozilla and individual contributors. See https://developer.mozilla.org/en-US/docs/MDN/About#Copyrights_and_licenses for details.

Ports can only be looked up using the 'ports' sub command. You can also look up ports inside a range. Information on ports comes from https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers under the Creative Commons Attribution-ShareAlike License.

I'm really interested in how you feel about this tool. Please take the time to fill in this short survey:
https://forms.gle/mHh6idwwUvdfYZM67
`),
	Run: root,
}

// Execute adds titles child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&titles, "titles", "t", titles, "List titles of the summaries available")
	rootCmd.PersistentFlags().IntVarP(&width, "width", "w", width, "Width to fit the output to")
	rootCmd.PersistentFlags().StringVarP(&searchTerm, "search", "", searchTerm, "Full-text search term")

	rootCmd.AddCommand(subCmd("html", "html", httpref.Html))
	rootCmd.AddCommand(subCmd("methods", "method", httpref.Methods))
	rootCmd.AddCommand(subCmd("statuses", "status", httpref.Statuses))
	rootCmd.AddCommand(subCmd("headers", "header", httpref.Headers))
	rootCmd.AddCommand(&cobra.Command{
		Use:     fmt.Sprintf("%s [filter]", "ports"),
		Aliases: []string{"port"},
		Short:   "References for common ports",
		Run:     portsReference(),
	})
}

func subCmd(name, alias string, ref httpref.References) *cobra.Command {
	return &cobra.Command{
		Use:     fmt.Sprintf("%s [filter]", name),
		Aliases: []string{alias},
		Short:   fmt.Sprintf("References for common HTTP %s", name),
		Run:     referenceCmd(ref),
	}
}

func root(cmd *cobra.Command, args []string) {
	if titles {
		results := append(httpref.Statuses.Titles(), httpref.Headers.Titles()...)
		results = append(results, httpref.Html.Titles()...)
		results = append(results, httpref.Methods.Titles()...)
		results = append(results, httpref.WellKnownPorts.Titles()...)
		results = append(results, httpref.RegisteredPorts.Titles()...)

		printResults(results)

		return
	}

	if searchTerm == "" && len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Must specify something to filter by\n")
		os.Exit(1)
	}

	results := append(httpref.Headers, httpref.Methods...)
	results = append(results, httpref.Statuses...)
	results = append(results, httpref.Html...)

	if searchTerm != "" {
		results = results.Search(searchTerm)
		printResults(results)

		return
	}

	results = results.ByName(args[0])
	printResults(results)
}

func printResults(results httpref.References) {
	switch len(results) {
	case 0:
		fmt.Fprintf(os.Stderr, "Filter not found any results\n")
		os.Exit(1)
	case 1:
		fmt.Printf("%s\n", results[0].Describe(width))
	default:
		for _, r := range results {
			fmt.Printf("%s\n", r.Summarize(width))
		}
	}
}

func referenceCmd(results httpref.References) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if searchTerm != "" {
			printResults(results.Search(searchTerm))
			return
		}

		if len(args) == 0 {
			printResults(results.Titles())
			return
		}

		printResults(results.ByName(args[0]))
	}
}

func portsReference() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if searchTerm != "" {
			fmt.Fprintf(os.Stderr, "error: full-text search not supported\n")
			os.Exit(1)
		}

		ref := append(httpref.WellKnownPorts, httpref.RegisteredPorts...)
		var results httpref.References

		if len(args) == 0 {
			results = ref.Titles()
		} else {
			results = ref.ByName(args[0])

			if len(results) == 0 {
				results = ref.InRange(args[0])
			}
		}

		printResults(results)
	}
}
