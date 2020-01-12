/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/dnnrly/httpref"
	"github.com/spf13/cobra"
)

// headersCmd represents the headers command
var headersCmd = &cobra.Command{
	Use:     "headers",
	Aliases: []string{"header"},
	Short:   "References for common HTTP headers",
	Long: `This displays useful information related to HTTP.

Most of the content comes from the Mozilla developer documentation (https://developer.mozilla.org/en-US/docs/Web/HTTP) and is copyright Mozilla and individual contributors. See https://developer.mozilla.org/en-US/docs/MDN/About#Copyrights_and_licenses for details.`,
	Run: func(cmd *cobra.Command, args []string) {
		results := httpref.Headers

		if len(args) == 0 {
			results = results.Titles()
		} else {
			results = results.ByName(args[0])
		}

		printResults(results)
	},
}

func init() {
	rootCmd.AddCommand(headersCmd)
}
