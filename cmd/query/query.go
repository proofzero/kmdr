/*
Copyright © 2021 kubelt

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
package query

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

var setupExtraHelp string = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}
`

var context string

// NewSetupCmd creates returns the apply command
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "run a query",
		Long:  `run a query`,
		RunE:  queryCmdRun,
	}

	cmd.Flags().StringVarP(&context, "context", "c", "", "select user context")

	cmd.SetHelpTemplate(setupExtraHelp)
	return cmd
}

// queryCmdRun prompts a query input
func queryCmdRun(cmd *cobra.Command, args []string) error {
	q := prompt.Input("> ", queryPrompt)
	fmt.Println("query:" + q)
	return nil
}

// Setup query prompt
func queryPrompt(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "user", Description: "Username"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
