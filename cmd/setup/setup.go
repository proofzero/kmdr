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
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
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
	API, err := api.NewAPI()
	if err != nil {
		return err
	}
	// generate keys for user
	pk, sk, err := util.GenerateUserKeys()
	if err != nil {
		p := &util.HelpPanic{
			Reason: "Error generating user keys: %s",
		}
		display, _ := p.Display(err)
		return errors.New(display)
	}
	// write the keypair out to users config directory
	// TODO: move to util or API and store keys in safer dir?
	home, _ := homedir.Dir()
	pkFile := fmt.Sprintf("%s/.config/kubelt/keys/%s.pub", home, username)
	skFile := fmt.Sprintf("%s/.config/kubelt/keys/%s", home, username)
	_, _ = os.Create(fmt.Sprintf("%s/.config/kubelt/keys", home)) // TODO: move this to a util
	// TODO: check errors
	_ = ioutil.WriteFile(pkFile, pk[:], 0644)
	_ = ioutil.WriteFile(skFile, sk[:], 0644)
	if err != nil {
		p := &util.HelpPanic{
			Reason: err.Error(),
		}
		display, _ := p.Display()
		return errors.New(display)
	}

	// create/update the config
	// TODO: handle errors
	configAPI := API.Config()
	_ = configAPI.InitConfig()
	_ = configAPI.AddContext("default", true)
	_ = configAPI.AddUser(username, true)
	_ = configAPI.Commit()

	// sync the new user to the cluster
	if err := API.Ktrl().IsAvailable(); err != nil {
		p := &util.HelpPanic{
			Reason: err.Error(),
		}
		display, _ := p.Display()
		return errors.New(display)
	}

	// Create the user
	cueAPI := API.Cue()
	user := make(map[string]string)
	user["data.name"] = username
	userVal, err := cueAPI.GenerateCueSpec("#user", user)
	if err != nil {
		p := util.HelpPanic{
			Reason: `ktrl is not running.`,
		}
		display, _ := p.Display()
		return errors.New(display)
	}

	_, err = API.Ktrl().Apply(userVal)
	if err != nil {
		p := util.HelpPanic{
			Reason: err.Error(),
		}
		display, _ := p.Display()
		return errors.New(display)
	}

	return nil
}
