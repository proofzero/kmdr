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
	"fmt"
	"io/ioutil"
	"os"

	"cuelang.org/go/cue"
	"github.com/proofzero/kmdr/pkg/kmdr"
	"github.com/proofzero/kmdr/pkg/ktrl"
	"github.com/proofzero/kmdr/pkg/util"
	"github.com/spf13/cobra"
)

var applyExtraHelp string = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}
`

var minFlagErr string = `Must supply a manifest with "-f" or "--filename"`

var file string

// NewContentCmd represents the content command
func NewApplyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "apply changes to kubelt resources",
		Long:  `apply changes to kubelt resources`,
		Args:  cobra.RangeArgs(0, 1),
		RunE:  applyCmdRun,
	}

	cmd.Flags().StringVarP(&file, "filename", "f", "", "object manifest")

	cmd.SetHelpTemplate(applyExtraHelp)
	return cmd
}

func applyCmdRun(cmd *cobra.Command, args []string) error {
	// TODO: migrate to a pre run?
	if file == "" {
		p := &util.HelpPanic{
			Reason: minFlagErr,
			Help:   applyExtraHelp,
			Error:  nil,
		}
		display, err := p.Display()
		if err != nil {
			return err
		}
		return errors.New(display)
	}

	var applyStr string
	if file == "-" {
		// Look at stdin for the good stuff
		reader := bufio.NewReader(os.Stdin)
		applyStr, _ = reader.ReadString('\n')
	} else { // read in cue file
		fBytes, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		applyStr = string(fBytes)
	}

	API := kmdr.NewAPI()
	validResources, err := runValidation(applyStr, API.Cue())
	if err != nil {
		return err
	}

	client, _ := ktrl.NewKtrlClient()
	err = applyResources(validResources, client)
	if err != nil {
		return err
	}

	return nil
}

func runValidation(applyStr string, cueAPI kmdr.CueAPI) (cue.Value, error) {
	applySchemas, err := cueAPI.CompileSchemaFromString(applyStr)
	if err != nil {
		return cue.Value{}, err
	}

	manifests := applySchemas.Value().LookupPath(cue.ParsePath("manifests"))
	schemas, _ := cueAPI.FetchSchema("kubelt://kmdr")
	mandef := schemas.LookupPath(cue.ParsePath("#manifests"))

	if err := cueAPI.ValidateResource(manifests, mandef); err != nil {
		return cue.Value{}, err
	}

	return applySchemas, err

	// TODO: migrate the below to ktrl
}

func applyResources(validResources cue.Value, client *ktrl.Client) error {
	resp, err := client.Apply(validResources)
	if err != nil {
		return err
	}
	for _, v := range resp.Resources.Cue {
		fmt.Println(v)
	}
	return nil
}
