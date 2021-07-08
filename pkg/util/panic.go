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
package util

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/fatih/color"
)

const UnhandledException = `An unhandled exception occurred.`

// HelpPanic
type HelpPanic struct {
	Reason string
	Help   string
	Error  error
}

// screen is the template for exporting meaningful errors to the users terminal
const screen = `
  {{define "Reason"}}  
  {{ .Reason }}
  {{end}}
  {{define "Help"}}
  If this problem persists, please open an issue at https://github.com/proofzero/kdmr/issues
  {{end}}
  {{define "Trace"}}
  TRACE:
  ---
  {{ .Error }}
  {{end}}`

// Display executes and returns the error templates with the provides errors
func (p *HelpPanic) Display(args ...interface{}) (string, error) {
	p.Reason = fmt.Sprintf(p.Reason, args...)
	tmpl, _ := template.New("error").Parse(screen)
	var str bytes.Buffer
	err := tmpl.ExecuteTemplate(&str, "Reason", p)
	if err != nil {
		return "", err
	}
	if p.Help != "" {
		err = tmpl.ExecuteTemplate(&str, "Help", p)
		if err != nil {
			return "", err
		}
	}
	if p.Error != nil {
		err = tmpl.ExecuteTemplate(&str, "Error", p)
		if err != nil {
			return "", err
		}
	}
	return color.RedString(str.String()), err
}
