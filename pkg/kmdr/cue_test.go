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
package kmdr

import (
	"fmt"
	"testing"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	gomock "github.com/golang/mock/gomock"
)

const validSchema string = `
manifests: {
	foo: {
		kind:       "Space"
		apiVersion: "kubelt.com/v1alpha1"
		metadata: {
			name: "test"
		}
	}
}
`

const invalidSchema string = `
manifests: {
	foo: {
		kind:       "Space"
		invalid: 	"error"
		metadata: {
			name: "test"
		}
	}
}
`

const validDefinition string = `
#manifests: [_=string]: {
    apiVersion: string
    kind: string
	metadata: {
		name: string
		[_]: _
	}
	data?: {
		[_]: _
	}
}
`

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
			args: args{
				apply: validSchema,
			},
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

func Test_cueAPI_ValidateResource(t *testing.T) {
	type fields struct{}
	type args struct {
		api   CueAPI
		apply string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "validates good schemas",
			fields: fields{},
			args: args{
				api: cueAPI{
					context: cuecontext.New(),
				},
				apply: validSchema,
			},
			wantErr: false,
		},
		{
			name:   "invalidates bad schemas",
			fields: fields{},
			args: args{
				api: cueAPI{
					context: cuecontext.New(),
				},
				apply: invalidSchema,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, _ := tt.args.api.CompileSchemaFromString(tt.args.apply)
			s, _ := tt.args.api.CompileSchemaFromString(validDefinition)
			val := v.LookupPath(cue.ParsePath("manifests"))
			def := s.LookupPath(cue.ParsePath("#manifests"))
			if err := tt.args.api.ValidateResource(val, def); (err != nil) != tt.wantErr {
				t.Errorf("cueAPI.ValidateResource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cueAPI_FetchSchema(t *testing.T) {
	type fields struct {
		context         *cue.Context
		schemasVersions *map[string]cue.Value
	}
	type args struct {
		apiVersion string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    cue.Value
		wantErr bool
	}{
		{
			name: "fetch schema",
			fields: fields{
				context: cuecontext.New(),
			},
			args: args{
				apiVersion: "kubelt://kmdr",
			},
			want:    cue.Value{},
			wantErr: false,
		},
		{
			name: "fail to fetch schema",
			fields: fields{
				context: cuecontext.New(),
			},
			args: args{
				apiVersion: "kubelt://invalid",
			},
			want:    cue.Value{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := NewMockCueAPI(ctrl)
			if tt.wantErr {
				invalidBytes := []byte("fesf:fei99{FEA")
				m.EXPECT().fetchSchema("kubelt://invalid").Return(invalidBytes, fmt.Errorf("foo"))
			} else {
				validBytes := []byte(validDefinition)
				m.EXPECT().fetchSchema("kubelt://kmdr").Return(validBytes, nil)
			}

			api := cueAPI{
				context:         tt.fields.context,
				schemasVersions: tt.fields.schemasVersions,
				schemaFetcher:   m.fetchSchema,
			}
			got, err := api.FetchSchema(tt.args.apiVersion)
			if got.Err() == nil && tt.wantErr {
				t.Errorf("cueAPI.FetchSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got.Err() != nil && !tt.wantErr {
				t.Errorf("cueAPI.FetchSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}
