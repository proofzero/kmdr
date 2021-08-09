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

import "testing"

func Test_help_Panic(t *testing.T) {
	type fields struct {
		Reason string
		Cta    string
		Err    error
	}
	type args struct {
		reason string
		cta    bool
		extra  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test panic w/ no cta and no extras",
			args: args{
				reason: "test",
				cta:    false,
				extra:  nil,
			},
			wantErr: false,
		},
		{
			name: "test panic w/ cta and no extras",
			args: args{
				reason: "test",
				cta:    true,
				extra:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHelp()
			if err := h.Panic(tt.args.reason, tt.args.cta, tt.args.extra...); (err == nil) != tt.wantErr {
				t.Errorf("help.Panic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
