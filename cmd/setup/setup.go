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
				help := util.Help(nil)
				return help.Panic("You must set --username or --anonymous flag when running setup.")
			} else if anon != "" {
				username = anon
			}
			return nil
		},
		RunE: setupCmdRun,
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "choose desired username")
	cmd.Flags().StringVarP(&anon, "anonymous", "a", "", "use an auto-generated username")
	cmd.Flags().BoolVarP(&ktrl, "init-ktrl", "k", true, "initialize an ktrl daemon if not already")

	cmd.SetHelpTemplate(setupExtraHelp)
	return cmd
}

// setupCmdRun bootstraps the local enviroment and configurations
func setupCmdRun(cmd *cobra.Command, args []string) error {
	help := util.Help(nil)
	API, err := api.NewAPI()
	if err != nil {
		return err
	}

	err = API.SetupUser(username)
	if err != nil {
		return help.Panic(err.Error())
	}

	return nil
}
