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
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/proofzero/kmdr/api"
	"github.com/proofzero/kmdr/util"
	"github.com/spf13/cobra"
)

var setupExtraHelp string = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}
`

var ktrl bool

// NewSetupCmd creates returns the apply command
func NewSetupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "setup kubelt kmdr",
		Long:  `setup kubelt kmdr`,
		RunE:  setupCmdRun,
	}

	cmd.Flags().BoolVarP(&ktrl, "init-ktrl", "k", true, "initialize an ktrl daemon if not already")

	cmd.SetHelpTemplate(setupExtraHelp)
	return cmd
}

// setupCmdRun bootstraps the local enviroment and configurations
func setupCmdRun(cmd *cobra.Command, args []string) error {
	help := util.NewHelp()
	API, err := api.NewAPI()
	if err != nil {
		return err
	}

	fmt.Println("Setting up kubelt user...")

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | green | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Please enter a username:",
		Templates: templates,
	}
	username, _ := prompt.Run()
	// TODO: add email validator for input
	prompt = promptui.Prompt{
		Label:     "Please enter your email address:",
		Templates: templates,
	}
	email, _ := prompt.Run()

	err = API.SetupUser(username, email)
	if err != nil {
		return help.Panic(err.Error(), true)
	}

	return nil
}
