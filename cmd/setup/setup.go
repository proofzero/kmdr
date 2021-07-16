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
	"errors"

	"github.com/proofzero/kmdr/api"
	"github.com/proofzero/kmdr/util"
	"github.com/spf13/cobra"
)

var setupExtraHelp string = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}
`

var (
	username string
	anon     string
	ktrl     bool
)

// NewSetupCmd creates returns the apply command
func NewSetupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "setup kubelt kmdr",
		Long:  `setup kubelt kmdr`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if username == "" && anon == "" {
				p := util.HelpPanic{
					Reason: "You must set --username or --anonymous flag when running setup.",
				}
				display, _ := p.Display()
				return errors.New(display)
			}
			return nil
		},
		RunE: setupCmdRun,
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "choose desired username")
	cmd.Flags().StringVarP(&anon, "anonymous", "a", "", "use an auto-generated username")
	cmd.Flags().BoolVarP(&ktrl, "init-ktrl", "k", true, "initialize an ktrl daemon if not already (default true)")

	cmd.SetHelpTemplate(setupExtraHelp)
	return cmd
}

// setupCmdRun bootstraps the local enviroment and configurations
func setupCmdRun(cmd *cobra.Command, args []string) error {
	// setup username keys

	API := api.NewAPI()

	if err := API.Ktrl().IsAvailable(); err != nil {
		return err
	}

	return nil
}
