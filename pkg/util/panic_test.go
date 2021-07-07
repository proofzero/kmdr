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
	"testing"
)

func TestHelpPanic_Display(t *testing.T) {
	type fields struct {
		Reason string
		Help   string
		Error  error
	}
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "generate error display",
			fields: fields{
				Reason: "test",
				Help:   "",
				Error:  nil,
			},
			args: args{},
		},
		// TODO: add tests for other sub templates and failures
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HelpPanic{
				Reason: tt.fields.Reason,
				Help:   tt.fields.Help,
				Error:  tt.fields.Error,
			}
			if got, err := p.Display(tt.args.args...); err != nil {
				t.Errorf("HelpPanic.Display() = \n%v", got)
			}
		})
	}
}
