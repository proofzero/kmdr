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
	"errors"
	"fmt"
	"text/template"

	"github.com/fatih/color"
)

// Interface
// -----------------------------------------------------------------------------

type Help interface {
	Panic(string, ...interface{}) error
}

// Implementation
// -----------------------------------------------------------------------------

const UnhandledException = `An unhandled exception occurred.`

// HelpPanic
type help struct {
	Reason string
	Cta    string
	Err    error
}

// screen is the template for exporting meaningful errors to the users terminal
const screen = `
  {{define "Reason"}}  
  {{ .Reason }}
  {{end}}
  {{define "Help"}}
  {{ .Cta }}
  {{end}}
  {{define "Trace"}}
  TRACE:
  ---
  {{ .Err }}
  {{end}}`

func NewHelp() *help {
	h := help{}
	h.Cta = "If this problem persists, please open an issue at https://github.com/proofzero/kdmr/issues"
	return &h
}

func (h *help) Panic(reason string, cta bool, extra ...interface{}) error {
	h.Reason = fmt.Sprintf(reason, extra...)

	tmpl, _ := template.New("error").Parse(screen)

	var str bytes.Buffer
	err := tmpl.ExecuteTemplate(&str, "Reason", h)
	if err != nil {
		return err
	}
	if cta {
		err = tmpl.ExecuteTemplate(&str, "Help", h)
		if err != nil {
			return err
		}
	}
	if h.Err != nil {
		err = tmpl.ExecuteTemplate(&str, "Error", h)
		if err != nil {
			return err
		}
	}
	return errors.New(color.RedString(str.String()))
}
