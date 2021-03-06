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
package version

import (
	"os"
	"strings"
	"text/template"

	_ "embed"

	"github.com/spf13/cobra"
)

// NewVersionCmd creates returns the version command
func NewVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Kmdr",
		Long:  `All software has versions. This is Kmdr's`,
		Run:   versionCmdRun,
	}
	return cmd
}

// versionCmdRun prints the version information to stdout
func versionCmdRun(cmd *cobra.Command, args []string) {
	PrintLogo()
}

// version represents a semver version format
type version struct {
	Major string
	Minor string
	Patch string
}

//go:embed version.txt
var versionTxt string

// The template for the version export
const kmdrASCII = `

██╗  ██╗███╗   ███╗██████╗ ██████╗
██║ ██╔╝████╗ ████║██╔══██╗██╔══██╗
█████╔╝ ██╔████╔██║██║  ██║██████╔╝	
██╔═██╗ ██║╚██╔╝██║██║  ██║██╔══██╗
██║  ██╗██║ ╚═╝ ██║██████╔╝██║  ██║
╚═╝  ╚═╝╚═╝     ╚═╝╚═════╝ ╚═╝  ╚═╝
v{{ .Major }}.{{ .Minor }}.{{ .Patch }}

`

// PrintLogo fetches current version information and executes a template
func PrintLogo() {
	semver := strings.Split(versionTxt, ".")
	v := version{Major: semver[0], Minor: semver[1], Patch: semver[2]}
	tmpl, _ := template.New("info").Parse(kmdrASCII)
	_ = tmpl.Execute(os.Stdout, v)
}
