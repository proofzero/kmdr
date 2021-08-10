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
package api

import (
	"reflect"
	"testing"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func Test_cueAPI_CompileSchemaFromString(t *testing.T) {
	type fields struct {
		context         *cue.Context
		schemasVersions *map[string]cue.Value
	}
	type args struct {
		apply string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    cue.Value
		wantErr bool
	}{
		{
			name: "generate schema from string",
			fields: fields{
				context: cuecontext.New(),
			},
			args:    args{},
			want:    cue.Value{},
			wantErr: false,
		},
		{
			name: "generate fails from invalid string",
			fields: fields{
				context: cuecontext.New(),
			},
			args: args{
				apply: "faesofh dfef",
			},
			want:    cue.Value{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := cueAPI{
				context:         tt.fields.context,
				schemasVersions: tt.fields.schemasVersions,
			}
			got, err := api.CompileSchemaFromString(tt.args.apply)
			if (err != nil) != tt.wantErr {
				t.Errorf("cueAPI.CompileSchemaFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Err() != nil && !tt.wantErr {
				t.Errorf("cueAPI.CompileSchemaFromString() = %v", got)
			}
		})
	}
}

func Test_cueAPI_unifyKmdrModule(t *testing.T) {
	type fields struct {
		context         *cue.Context
		definitions     cue.Value
		schemasVersions *map[string]cue.Value
		schemaFetcher   func(apiVersion string) ([]byte, error)
	}
	type args struct {
		input cue.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   cue.Value
	}{
		// TODO: Add test cases for unifying cue.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := cueAPI{
				context:         tt.fields.context,
				definitions:     tt.fields.definitions,
				schemasVersions: tt.fields.schemasVersions,
				schemaFetcher:   tt.fields.schemaFetcher,
			}
			if got := api.unifyKmdrModule(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cueAPI.unifyKmdrModule() = %v, want %v", got, tt.want)
			}
		})
	}
}
