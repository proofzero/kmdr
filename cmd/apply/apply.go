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
package apply

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"

	"github.com/proofzero/kmdr/api"
	"github.com/proofzero/kmdr/util"
	"github.com/spf13/cobra"
)

var applyExtraHelp string = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}
`

var file string

// NewApplyCmd creates returns the apply command
func NewApplyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "apply changes to kubelt resources",
		Long:  `apply changes to kubelt resources`,
		Args:  cobra.RangeArgs(0, 1),
		RunE:  applyCmdRun,
	}

	cmd.Flags().StringVarP(&file, "filename", "f", "", "object manifest")
	_ = cmd.MarkFlagRequired("filename")

	cmd.SetHelpTemplate(applyExtraHelp)
	return cmd
}

// applyCmdRun accepts stdnin or a cue file, validates the contents
// before sending the contents to ktrl
func applyCmdRun(cmd *cobra.Command, args []string) error {
	var applyStr string
	if file == "-" {
		// Look at stdin for the good stuff
		reader := bufio.NewReader(os.Stdin)
		applyStr, _ = reader.ReadString('\x1D')
	} else { // read in cue file
		fBytes, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		applyStr = string(fBytes)
	}

	API, _ := api.NewAPI()
	err := API.Apply(applyStr)
	if err != nil {
		p := &util.HelpPanic{
			Reason: err.Error(),
		}
		display, _ := p.Display()
		return errors.New(display)
	}

	return nil
}
