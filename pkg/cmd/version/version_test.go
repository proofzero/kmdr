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
	"testing"

	"github.com/spf13/cobra"
)

func Test_versionCmdRun(t *testing.T) {
	cmd := NewVersionCmd()
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "successfully print version",
			args: args{cmd: cmd},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd.Run(tt.args.cmd, tt.args.args)
		})
	}
}

func TestPrintLogo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "successfully print logo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintLogo()
		})
	}
}
