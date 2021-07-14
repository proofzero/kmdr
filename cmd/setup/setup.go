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
package setup

import (
	"github.com/spf13/cobra"
)

var setupExtraHelp string = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}
`

// NewApplyCmd creates returns the apply command
func NewSetupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "setup kubelt kmdr",
		Long:  `setup kubelt kmdr`,
		RunE:  setupCmdRun,
	}

	// cmd.Flags().StringVarP(&file, "filename", "f", "", "object manifest")
	// cmd.MarkFlagRequired("filename")

	cmd.SetHelpTemplate(setupExtraHelp)
	return cmd
}

// setupCmdRun accepts stdnin or a cue file, validates the contents
// before sending the contents to ktrl
func setupCmdRun(cmd *cobra.Command, args []string) error {
	// client, _ := api.NewKtrlClient()
	return nil
}